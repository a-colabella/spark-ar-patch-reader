/*
  This go program parses Spark AR patch json data
*/

package arputil

// Gets the name of the patch
func GetName(patch map[string]interface{}) string {
  children := patch["children"].([]interface{})
  first_child := children[0].(map[string]interface{})
  name := first_child["name"].(string)
  return name
}

// Returns the number of input ports in the group patch
func CountInputPorts(patch map[string]interface{}) int {
  children := patch["children"].([]interface{})
  first_child := children[0].(map[string]interface{})
  grandchildren := first_child["children"].([]interface{})

  var input_ports map[string]interface{}

  for _, grandchild := range grandchildren {
    if grandchild.(map[string]interface{})["modelName"] == "patch_input_ports_model" {
      input_ports = grandchild.(map[string]interface{})
    }
  }

  input_children := input_ports["children"].([]interface{})

  return len(input_children)
}

// Returns the number of patches in the group patch
func CountPatches(patch map[string]interface{}) int {
  children := patch["children"].([]interface{})
  first_child := children[0].(map[string]interface{})
  grandchildren := first_child["children"].([]interface{})

  var patches map[string]interface{}

  for _, grandchild := range grandchildren {
    if grandchild.(map[string]interface{})["modelName"] == "patches_model" {
      patches = grandchild.(map[string]interface{})
    }
  }

  patch_children := patches["children"].([]interface{})

  return len(patch_children)
}

// Returns the number of connections in the group patch
func CountConnections(patch map[string]interface{}) int {
  children := patch["children"].([]interface{})
  first_child := children[0].(map[string]interface{})
  grandchildren := first_child["children"].([]interface{})

  var connections map[string]interface{}

  for _, grandchild := range grandchildren {
    if grandchild.(map[string]interface{})["modelName"] == "connections_model" {
      connections = grandchild.(map[string]interface{})
    }
  }

  connection_children := connections["children"].([]interface{})

  return len(connection_children)
}

// Returns the number of output ports in the group patch
func CountOutputPorts(patch map[string]interface{}) int {
  children := patch["children"].([]interface{})
  first_child := children[0].(map[string]interface{})
  grandchildren := first_child["children"].([]interface{})

  var output_ports map[string]interface{}

  for _, grandchild := range grandchildren {
    if grandchild.(map[string]interface{})["modelName"] == "patch_output_ports_model" {
      output_ports = grandchild.(map[string]interface{})
    }
  }

  output_children := output_ports["children"].([]interface{})

  return len(output_children)
}
