#!/usr/bin/env ruby
require 'rubygems'
require 'mqtt'
$stdout.sync=true
puts 'MQTT ruby publisher'
c = MQTT::Client.connect('localhost')
i = 0
while true
  c.publish 'my_topic','rubyist: '+i.to_s  # publish // HL
  i = i+1
  sleep(1)
end
