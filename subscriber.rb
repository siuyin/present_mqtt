#!/usr/bin/env ruby
require 'rubygems'
require 'mqtt'
$stdout.sync = true
puts "Ruby MQTT subscriber"
c = MQTT::Client.connect 'localhost'
c.get('my_topic') do |topic, msg|   # subscribe // HL
  puts msg
end
