package config

import "strings"

var ExampleYaml = `
devices:
  light.kitchen:
    group: downstairs
    name: Kitchen
    type: light
endpoints:
  nanomsg:
    pub: tcp://127.0.0.1:8792
    sub: tcp://127.0.0.1:8793
  mqtt:
    broker: tcp://127.0.0.1:1883
processes:
  nonexistent:
    cmd: gohome service nonexistent
bill:
  electricity:
    primary_rate: 8.54
    standing_charge: 18.9
  gas:
    calorific_value: 39.3
    conversion_factor: 1.02264
    multiplier: 0.01
    primary_rate: 3.16
    standing_charge: 35.4
  vat: 0
  currency: £
earth:
  latitude: 51.5072
  longitude: 0.1275
espeak:
  args: -s 140
general:
  email:
    admin: admin@example.com
    from: me@example.com
    server: localhost:25
heating:
  device: heater.boiler
  schedule:
    hallway:
      Monday,Tuesday,Wednesday,Thursday,Friday:
      - '0:00': 0
      Saturday,Sunday:
      - '9:00': 15.5
      - '22:30': 10
    living:
      Friday:
      - '7:45': 16
      - '8:10': 14
      - '16:45': 17
      - '22:20': 10
      Monday,Tuesday,Wednesday,Thursday:
      - '7:40': 16.5
      - '8:10': 14
      - '17:30': 17
      - '22:15': 10
      Saturday,Sunday:
      - '9:00': 16
      - '22:50': 10
    office:
      Monday,Tuesday,Wednesday,Thursday,Friday:
      - '9:00': 16
      - '18:00': 10
      Saturday,Sunday:
      - '0:00': 0
    unoccupied:
      Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday:
      - '0:00': 6
  sensors:
  - temp.hallway
  - temp.living
  - temp.office
  slop: 0.1
irrigation:
  at: 6h
  device: pump.garden
  enabled: true
  factor: 1.5
  interval: 12h
  max_temp: 25
  max_time: 60
  min_temp: 13
  min_time: 10
jabber:
  jid: myjabberid@gmail.com/gohome
  pass: password
protocols:
  arduino:
    A: lock.front
    C: alarm.bell
  chime:
    byronsx.0001: bell.front
  earth:
    home: earth
  energenie:
    00097f: trv.living
    00098b: trv.spareroom
  heating:
    ch: heating.ch
  homeeasy:
    07393AFA: door.front
    0393BDEA: letterbox.front
    00123451: light.bedside
    00123452: light.hallway
    00123453: light.glowworm
    00123454: heater.blanket
    00123455: heater.boiler
    00123456: amp.living
    00123458: pump.garden
    00234561: light.bedside
    00234564: heater.blanket
    0493BD1A: pir.kitchen
    09FD62BA: pir.living
  power:
    cc01: power.power
  pressure:
    wmr100: pressure.pressure
  rain:
    wmr100.0: rain.outside
  rfid:
    "0000000001": rfid.person1
    "0000000002": rfid.person2
  temp:
    thn132n.12: temp.outside
    thn132n.34: temp.living
    wmr100.0: temp.office
    wmr100.1: temp.garden
    wmr100.2: temp.hallway
  wind:
    wmr100: wind.outside
  x10:
    b01: light.bedside
    b02: heater.blanket
    b03: light.glowworm
    b04: pump.garden
    b06: light.kitchen
    b07: pir.upstairs
    m01: pir.front
    m02: pir.side
    m03: pir.back
    m08: pir.garage
  xpl:
    slimdev-slimserv.FrontroomTouch: hifi.living
    slimdev-slimserv.KitchenRadio: hifi.kitchen
sms:
  telephone: '+441234567890'
twitter:
  auth:
    consumer_key: xxx
    consumer_secret: yyy
    token: aaa
    token_secret: bbb
voice:
  'lights? on':
    switch glowworm on
  'lights? off':
    switch glowworm off
  'blanket on':
    switch blanket on
  'blanket off':
    switch blanket off
watchdog:
  devices:
    power.power: 5m
    temp.garden: 20m
    temp.hallway: 4h
    temp.living: 20m
    temp.office: 20m
    temp.outside: 20m
    wind.outside: 20m
weather:
  outside:
    rain: rain.outside
    temp: temp.garden
    wind: wind.outside
  windy: 3.2
wunderground:
  id: STATIONID
  password: password
  url: http://weatherstation.wunderground.com/weatherstation/updateweatherstation.php`

var ExampleConfig, _ = OpenReader(strings.NewReader(ExampleYaml))
