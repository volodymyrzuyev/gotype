# Use Cases
- [x] 1. Be able to view and scroll thrugh a text file
- [x] 2. Have simular navigation as vim
    - [x] up/down/left/righ - k/j/h/l
    - w to move up 1 word
    - b to move down 1 word
- [x] 3. Able to save a text file to disk
- [?] 4. Be responsive nad not feel "slow" on my computer
- [x] 5. implement modes 
    - "normal" : where you scroll thrugh text
    - "editing" : where keys you press are typed into the document
- [ ] 6. implement undo - press a key and go back 1 "editing" sesion

## Implementation
1. Have a screen which displays a part of the document that will fit into the current terminal window.
    - lines should not wrap. If a line is longer than the screen, than scroll the whole text so the cursor will fit on screen
2. self explementory, press key, cursol move
3. Save file when a command is typed in
4. Don't be slow
5. In noraml move, keys are commands, in editing mode, keys are text
6. save couple of states, so that stuff can be undone, probably a file on disk
