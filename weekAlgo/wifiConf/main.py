import wifi
import time
import subprocess

def allWindows(cls, interface):
        """
        Returns a list of all cells extracted from the output of iwlist.
        """
        #try:
            #iwlist_scan = subprocess.check_output(["netsh", interface, "show","network"])
        #except subprocess.CalledProcessError as e:
        #    raise InterfaceError(e.output.strip())
        #else:
        #    iwlist_scan = iwlist_scan.decode('utf-8')
        cells = map(Cell.from_string, cells_re.split(iwlist_scan)[1:])

        return cells

#available_networks = wifi.Cell.allWindows("wlan")
wifi_networks = wifi.Cell.allWindows("wlan")

#available_networks = subprocess.check_output(["netsh", "wlan","show","network"])
#available_networks = available_networks.decode('utf-8')

print(wifi_networks)


byteList = list(wifi_networks)

newSentence = ""
newList = list()

for char in byteList:
    if char == "\n":
        newList.append(newSentence)
        newSentence = ""
        continue
    if char == "\r":
        continue
    else:
        newSentence += char
#for sentence in newList:
#     if "SSID" in sentence:
#        print(sentence)

print(*[sentence for sentence in newList if "SSID" in sentence], sep="\n")
print(wifi_networks)
