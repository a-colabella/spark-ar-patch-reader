/*
  This go program parses Spark AR patch json data
*/

package arputil

import "strconv"

func patchModel(patch map[string]interface{}, property string) string {
  children := patch["children"].([]interface{})
  first_child := children[0].(map[string]interface{})
  model_name := ""

  switch property {
  case "name":
    return first_child["name"].(string)
  case "input":
    model_name = "patch_input_ports_model"
  case "output":
    model_name = "patch_output_ports_model"
  case "patches":
    model_name = "patches_model"
  case "connections":
    model_name = "connections_model"
  }

  grandchildren := first_child["children"].([]interface{})
  var xs map[string]interface{}

  for _, grandchild := range grandchildren {
    if grandchild.(map[string]interface{})["modelName"] == model_name {
      xs = grandchild.(map[string]interface{})
    }
  }

  ys := xs["children"].([]interface{})

  return strconv.Itoa(len(ys))
}

// Gets the name of the patch
func GetName(patch map[string]interface{}) string {
  return patchModel(patch, "name")
}

// Returns the number of input ports in the group patch
func CountInputPorts(patch map[string]interface{}) string {
  return patchModel(patch, "input")
}

// Returns the number of patches in the group patch
func CountPatches(patch map[string]interface{}) string {
  return patchModel(patch, "patches")
}

// Returns the number of connections in the group patch
func CountConnections(patch map[string]interface{}) string {
  return patchModel(patch, "connections")
}

// Returns the number of output ports in the group patch
func CountOutputPorts(patch map[string]interface{}) string {
  return patchModel(patch, "output")
}
