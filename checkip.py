import socket

from datetime import datetime
import time
import winsound
from termcolor import colored

sos = "...---..."
MAKE_BEEP = True

def getip():
    hostname = socket.gethostname()
    local_ip = socket.gethostbyname(hostname)
    return local_ip

def makebeep():
	for i in range(len(sos)):
		if sos[i] == ".":
			winsound.Beep(1000, 200)
		else: winsound.Beep(1000, 400)

if __name__ == '__main__':
	ip = getip()
	dtime = datetime.now()
	print(ip + " ("+ str(dtime) +")")
	while True:
		newip = getip()
		if newip != ip:
			ip = newip
			if ip == "127.0.0.1":
				color = "yellow"
			else:
				color = "red"
			if MAKE_BEEP: makebeep()
			print(colored("    " + str(datetime.now()), "yellow")+colored(" -> "+ip, color))
		time.sleep(10)