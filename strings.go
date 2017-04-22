package main

const MsgWelcome = `
Welcome to the Maru installer!

In order to install Maru you will need to:

1. Connect your device to your computer over USB

2. Enable USB Debugging on your device:

    1)  Go to the Settings app and scroll down to
        the System section

        NOTE: If you already have "Developer options"
        under System then go directly to #5

    2)  Tap on "About phone"
    3)  Tap "Build number" 7 times until you get a message
        that says you are now a developer
    4)  Go back to the main Settings app
    5)  Tap on "Developer options"
    6)  Ensure that "USB debugging" is enabled
    7)  Tap "OK" if you see a dialog asking you to allow
        USB Debugging for your computer's RSA key fingerprint

WARNING: Installing Maru will wipe all your personal data
so make sure you first back-up anything important!
`

const MsgIncompleteZip = `
Hmm, looks like your installer is a missing a few things.

Are you running this install script outside the directory you unzipped the installer in?
`

const msgFixAdb = `
Please ensure that:

1. Your device is connected to your computer over USB
2. You have USB Debugging enabled (see above for instructions)
3. You unlock your device and tap "OK" if you see a dialog asking you
   to allow USB Debugging for your computer's RSA key fingerprint

Go ahead and re-run the installer when you're ready.
`
const MsgAdbIssue = "\nHmm, there was an issue communicating with your device.\n" + msgFixAdb

const MsgFastbootNoDeviceFound = `
Hmm, your device can't be found. Please ensure that your device is connected to your computer over USB.
`

const MsgFixPerms = `
It looks like you are missing some USB permissions.

On certain Linux distributions (Ubuntu 14.04 for example), you will need to
explicitly add permissions to access USB devices:

1. Disconnect your device from USB

2. Run this in a terminal (requires sudo):

   $ wget -S -O - https://source.android.com/source/51-android.txt | sed "s/<username>/$USER/" | sudo tee >/dev/null /etc/udev/rules.d/51-android.rules; sudo udevadm control --reload-rules

3. Re-connect your device over USB and re-run this installer
`

const MsgUnlockSuccess = `
Successfully unlocked bootloader!

Your device will need to reboot before continuing. It will factory reset, so
this reboot can take a few minutes longer than usual.

To continue the installation process, please re-run this script after your
device completely boots up and you have re-enabled USB Debugging.
`

const MsgSuccess = `
Installation complete!

The first boot will take 2-3 mins as Maru sets up your device so please be
patient.

Rebooting into Maru...
`
