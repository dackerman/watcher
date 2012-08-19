Simple directory watcher program in Go.
=======================================

Run from command line
---------------------
    go get github.com/dackerman/watcher
    folderwatcher -folder="~/mychangingdir" -cmd="sh myshellcommand.sh" -recurse=true

Invoke from Go Code
-------------------
    // notify is a read-only chan that is pinged when the folder changes
    notify := watcher.WatchDirectory("~/mychangingdir", recurse)

    watcher.ExecuteOnChange(notify, func() {
        // This gets executed when the folder changes
    })