package resourcecollector

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	stork_api "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// Since we collect all resources from required migration namespace at once
// getResourcePatch creates map of namespace: {kind: []resourceinfo{}}
// to get transform spec for matching resources
func GetResourcePatch(transformName string, namespaces []string) (map[string]stork_api.KindResourceTransform, error) {
	// namespace- Kind:TransformSpec map for faster lookup
	patch := make(map[string]stork_api.KindResourceTransform)
	if transformName == "" {
		logrus.Error("Empty name received for resource transformation")
		return patch, nil
	}
	for _, namespace := range namespaces {
		resp, err := storkops.Instance().GetResourceTransformation(transformName, namespace)
		if err != nil {
			// current namespace does not have any transform CR
			// skip it from map
			if errors.IsNotFound(err) {
				continue
			}
			logrus.Errorf("Unable to get resource transfomration specs %s/%s, err: %v", namespace, transformName, err)
			return nil, err
		}
		resMap := make(map[string][]stork_api.TransformResourceInfo)
		for _, resource := range resp.Status.Resources {
			resMap[resource.Kind] = append(resMap[resource.Kind], *resource)
		}
		patch[namespace] = resMap
	}
	return patch, nil
}

// this method transform object as per resource transformation specified in each namespaces
func TransformResources(
	object runtime.Unstructured,
	resPatch []stork_api.TransformResourceInfo,
	objName, objNamespace string,
) error {
	for _, patch := range resPatch {
		if patch.Name == objName && patch.Namespace == objNamespace {
			content := object.UnstructuredContent()
			for _, path := range patch.Specs.Paths {
				switch path.Operation {
				case stork_api.AddResourcePath:
					value, err := getNewValueForPath(path.Value, path.Type)
					if err != nil {
						logrus.Errorf("Unable to parse the Value for the type %s specified, path %s on resource kind: %s/,%s/%s,  err: %v", path.Type, path, patch.Kind, patch.Namespace, patch.Name, err)
						return err
					}
					if path.Type == stork_api.KeyPairResourceType {
						updateMap := value.(map[string]string)
						err := SetNestedStringMap(content, updateMap, path.Path)
						if err != nil {
							logrus.Errorf("Unable to apply patch path %s on resource kind: %s/,%s/%s,  err: %v", path, patch.Kind, patch.Namespace, patch.Name, err)
							return err
						}
					} else {
						err := SetNestedField(content, value, path.Path)
						if err != nil {
							logrus.Errorf("Unable to perform operation %s on path %s on resource kind: %s/,%s/%s,  err: %v", path.Operation, path, patch.Kind, patch.Namespace, patch.Name, err)
							return err
						}
					}

				case stork_api.DeleteResourcePath:
					RemoveNestedField(content, strings.Split(path.Path, ".")...)
					logrus.Debugf("Removed patch path %s on resource kind: %s/,%s/%s", path, patch.Kind, patch.Namespace, patch.Name)

				case stork_api.ModifyResourcePathValue:
					var value interface{}
					if path.Type == stork_api.KeyPairResourceType {
						currMap, _, err := unstructured.NestedMap(content, strings.Split(path.Path, ".")...)
						if err != nil || len(currMap) == 0 {
							return fmt.Errorf("unable to find spec path, err: %v", err)
						}
						mapList := strings.Split(path.Value, ",")
						for _, val := range mapList {
							keyPair := strings.Split(val, ":")
							if len(keyPair) != 2 {
								return fmt.Errorf("invalid keypair value format :%s", keyPair)
							}
							currMap[keyPair[0]] = keyPair[1]
						}
						value = currMap
					} else if path.Type == stork_api.SliceResourceType {
						currList, _, err := unstructured.NestedSlice(content, strings.Split(path.Path, ".")...)
						if err != nil {
							return fmt.Errorf("unable to find spec path, err: %v", err)
						}
						arrList := strings.Split(path.Value, ",")
						for _, val := range arrList {
							currList = append(currList, val)
						}
						value = currList
					} else {
						var err error
						value, err = getNewValueForPath(path.Value, path.Type)
						if err != nil {
							logrus.Errorf("Unable to parse the Value for the type %s specified, path %s on resource kind: %s/,%s/%s,  err: %v", path.Type, path, patch.Kind, patch.Namespace, patch.Name, err)
							return err
						}
					}
					err := SetNestedField(content, value, path.Path)
					if err != nil {
						logrus.Errorf("Unable to perform operation %s on path %s on resource kind: %s/,%s/%s,  err: %v", path.Operation, path, patch.Kind, patch.Namespace, patch.Name, err)
						return err
					}
				}
			}
			// lets add annotation saying this resource has been transformed by migration/restore
			// controller before applying
			// set migration annotations
			annotations, found, err := unstructured.NestedStringMap(content, "metadata", "annotations")
			if err != nil {
				return err
			}
			if !found {
				annotations = make(map[string]string)
			}
			annotations[TransformedResourceName] = "true"
			if err := unstructured.SetNestedStringMap(content, annotations, "metadata", "annotations"); err != nil {
				return err
			}
			object.SetUnstructuredContent(content)
			logrus.Infof("Updated resource of kind %v with patch , resource: %v", patch.Kind, object)
		}
	}
	return nil
}

func getNewValueForPath(oldVal string, valType stork_api.ResourceTransformationValueType) (interface{}, error) {
	var updatedValue interface{}
	var err error

	switch valType {
	case stork_api.KeyPairResourceType:
		//TODO: here we can accept map[string]interface{} because inside it is getting changed
		newVal := make(map[string]string)
		mapList := strings.Split(oldVal, ",")
		for _, val := range mapList {
			keyPair := strings.Split(val, ":")
			newVal[keyPair[0]] = keyPair[1]
		}
		updatedValue = newVal
	case stork_api.SliceResourceType:
		newVal := []string{}
		arrList := strings.Split(oldVal, ",")
		newVal = append(newVal, arrList...)
		updatedValue = newVal
	case stork_api.IntResourceType:
		updatedValue, err = strconv.ParseInt(oldVal, 10, 64)
	case stork_api.BoolResourceType:
		updatedValue, err = strconv.ParseBool(oldVal)
	}
	return updatedValue, err
}

func jsonPath(fields []string) string {
	return "." + strings.Join(fields, ".")
}

var pathRegexpWithanArray = regexp.MustCompile(`^.+\[[0-9]+\](\.[a-zA-Z_/][a-zA-Z0-9_/]*)+$`)

func RemoveNestedField(obj map[string]interface{}, fields ...string) {
	if !pathRegexpWithanArray.MatchString(strings.Join(fields, ".")) {
		unstructured.RemoveNestedField(obj, fields...)
		return
	}
	m := obj
	for _, field := range fields[:len(fields)-1] {
		if val, ok := getMapKeyValue(m, field); ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return
			}
		} else {
			return
		}
	}
	delete(m, fields[len(fields)-1])
}

func SetNestedStringMap(obj map[string]interface{}, value map[string]string, path string) error {
	if !pathRegexpWithanArray.MatchString(path) {
		return unstructured.SetNestedStringMap(obj, value, strings.Split(path, ".")...)
	}
	m := make(map[string]interface{}, len(value)) // convert map[string]string into map[string]interface{}
	for k, v := range value {
		m[k] = v
	}
	return setNestedFieldNoCopy(obj, m, strings.Split(path, ".")...)
}

func SetNestedField(obj map[string]interface{}, value interface{}, path string) error {
	if !pathRegexpWithanArray.MatchString(path) {
		return unstructured.SetNestedField(obj, value, strings.Split(path, ".")...)
	}
	return setNestedFieldNoCopy(obj, runtime.DeepCopyJSONValue(value), strings.Split(path, ".")...)
}

func setNestedFieldNoCopy(obj map[string]interface{}, value interface{}, fields ...string) error {
	m := obj

	for i, field := range fields[:len(fields)-1] {
		if val, ok := getMapKeyValue(m, field); ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", jsonPath(fields[:i+1]))
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value
	return nil
}

func getMapKeyValue(m map[string]interface{}, field string) (interface{}, bool) {
	//TODO: Can use regexp.MustCompile("")
	f := func(c rune) bool {
		return c == '[' || c == ']'
	}
	parts := strings.FieldsFunc(field, f)
	if len(parts) != 2 {
		value, ok := m[field]
		return value, ok
	}
	arr := m[parts[0]]
	value, ok := arr.([]interface{})
	if !ok {
		return m[parts[0]], true
	}

	var index int
	fmt.Sscanf(parts[1], "%d", &index)
	if index < len(value) {
		return value[index], true
	} else if index == len(value) {
		// try creating another func use def - it may create even during RemoveNestedField
		value = append(value, make(map[string]interface{}))
		return value[index], true
	}
	logrus.Errorf("index [%s] is beyound the array: %s with length %d", parts[1], parts[0], len(value))
	// TODO: Test this case
	return nil, false
}

// func def(m map[string]interface{}, field string) map[string]interface{} {
// 	newVal := make(map[string]interface{})

// 	f := func(c rune) bool {
// 		return c == '[' || c == ']'
// 	}
// 	parts := strings.FieldsFunc(field, f)
// 	if len(parts) != 2 {
// 		m[field] = newVal
// 		return m
// 	}
// }
