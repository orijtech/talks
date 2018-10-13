func waitForIt(ch <-chan os.Signal) {
        sig := <-ch

        // After we've received the signal log it then exit
        // ...
}
