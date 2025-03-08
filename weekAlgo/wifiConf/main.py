import wifi
import time

wifi_scanner = wifi.Cell.all("wlan")
available_networks = [print(cell.ssid) for cell in wifi_scanner]