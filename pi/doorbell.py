import RPi.GPIO as GPIO
import time, os
from subprocess import call

gpio_pin = 18
script_path = os.path.join(os.path.dirname(__file__), 'notify.sh')

def cb(pin):
    print("button pressed")
    sys.stdout.flush()
    call([script_path])

GPIO.setmode(GPIO.BCM)
GPIO.setup(gpio_pin, GPIO.IN, pull_up_down=GPIO.PUD_DOWN)
GPIO.add_event_detect(gpio_pin, GPIO.RISING, callback=cb)

while True:
    next

GPIO.cleanup()
