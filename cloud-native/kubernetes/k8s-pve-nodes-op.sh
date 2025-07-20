#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

main() { 
  local action="start"
  
  # Parse command line arguments
  while [[ $# -gt 0 ]]; do
    case $1 in
      --action=*)
        action="${1#*=}"
        shift
        ;;
      *)
        echo "Unknown argument: $1"
        exit 1
        ;;
    esac
  done

  # Validate action
  if [[ ! "$action" =~ ^(start|stop|shutdown)$ ]]; then
    echo "Error: Action must be 'start' or 'stop' or 'shutdown'"
    exit 2
  fi

  # Define VM IDs per node (using a list of node-VM pairs)
  local -a vm_list=(
    "m7:5210"
    "m7:5310"
    "ser6:7000"
    "ser6:7310"
    "gem12:9000"
    "gem12:9310"
  )

  # Process each VM in the list
  for node_vm in "${vm_list[@]}"; do
    local node="${node_vm%%:*}"
    local vmid="${node_vm##*:}"
    
    echo "${action}ing VM $vmid on node $node"
    
    # Check if node is reachable
    if ! ping -c 1 -W 1 "$node" &> /dev/null; then
      echo "Error: Unable to reach node $node"
      continue
    fi
    
    # Using SSH to connect to PVE node and execute qm command
    if ! ssh "root@$node" "qm ${action} $vmid"; then
      echo "Error: Failed to ${action} VM $vmid on node $node"
    else
      echo "Successfully ${action}ed VM $vmid on node $node"
    fi
  done
}

main "$@"