package ruby

// http://sandbox.mc.edu/~bennet/ruby/code/ctest3_rb.html

const EightBitAdder = `
#
# This is a one-bit adder.
#

require "csim"
require "cgrp"
require "oba"

NumberOut.shush

# Blueprint for a the one-bit adder
bp = OBA.new

# Two input senders, and the output device.
na = SwitchBank.new
nb = SwitchBank.new
disp = NumberOut.new("  Sum")

# We're going to build an 8-bit adder
prev = nil
8.times do
  # Create the one-bit adder and join the data inputs and outputs.
  addr = bp.another
  na.join(addr)
  nb.join(addr)
  addr.join(disp)

  # Chain the carry, if this isn't he first one.
  if prev then
    prev.join(addr)
  end

  prev = addr
end

# Overflow light.
prev.join(LED.new("  Oflow"))

NumberOut.shush(false)

30.times do
  a = rand(256)
  b = rand(256)
  print a, " + ", b, ":\n"
  Gate.activate
  na.value = a
  nb.value = b
  Gate.deactivate
end
`
