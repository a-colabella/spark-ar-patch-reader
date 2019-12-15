/*
  This go program parses Spark AR patch json data
*/

package arputil

import "strconv"

// Accepts an input patch and a property string
// and returns the corresponding name or count as string
func patchModel(patch Patch, property string) string {
  var result string
  var xs int

  if property == "name" {
    result = patch.Children[0].Name
  } else {
    grandchildren := patch.Children[0].Children

    for _, grandchild := range grandchildren {
      if grandchild.ModelName == property {
        xs = len(grandchild.Children)
      }
    }

    result = strconv.Itoa(xs)
  }

  return result
}

// Gets the name of the patch
func GetName(patch Patch) string {
  return patchModel(patch, "name")
}

// Returns the number of input ports in the group patch
func CountInputPorts(patch Patch) string {
  return patchModel(patch, "patch_input_ports_model")
}

// Returns the number of patches in the group patch
func CountPatches(patch Patch) string {
  return patchModel(patch, "patches_model")
}

// Returns the number of connections in the group patch
func CountConnections(patch Patch) string {
  return patchModel(patch, "connections_model")
}

// Returns the number of output ports in the group patch
func CountOutputPorts(patch Patch) string {
  return patchModel(patch, "patch_output_ports_model")
}
