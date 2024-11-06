"""Cluster profile of m510 instances with f=1

Instructions:
check BFTBrain's `README.md`
"""

#
# NOTE: This code was machine converted. An actual human would not
#       write code like this!
# Edited by Haoyun Qin

fault_count = 1
hardware_type = "m510"

node_count = 3 * fault_count + 1 + 2

# Import the Portal object.
import geni.portal as portal

# Import the ProtoGENI library.
import geni.rspec.pg as pg

# Import the Emulab specific extensions.
import geni.rspec.emulab as emulab

# Create a portal object,
pc = portal.Context()

# Create a Request object to start building the RSpec.
request = pc.makeRequestRSpec()

ifaces = []

for i in range(node_count):
    node = request.RawPC("node-" + str(i))
    node.hardware_type = hardware_type
    iface = node.addInterface("interface-" + str(i))
    ifaces.append(iface)

# Link link-1
link_1 = request.Link("link-1")
link_1.Site("undefined")

for iface in ifaces:
    link_1.addInterface(iface)

# Print the generated rspec
pc.printRequestRSpec(request)
