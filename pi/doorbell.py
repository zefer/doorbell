import RPi.GPIO as GPIO
import time, os
from subprocess import call

GPIO.setmode(GPIO.BCM)

GPIO.setup(18, GPIO.IN, pull_up_down=GPIO.PUD_UP)

while True:
    input_state = GPIO.input(18)
    if input_state == False:
        print('Button Pressed')
        script_path = os.path.join(os.path.dirname(__file__), 'notify.sh')
        call([script_path])
        time.sleep(0.5)
