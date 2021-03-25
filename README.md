# fasttrack-qconsumer

This is the CLI to the qserver, which needs to be started beforehand.

There are two simple commands :

<code>qconsumer qget</code> Which displays a list of fiver generic questions on the screen each with a unique ID and three options like such

Question no [74]<br>
 The Question<br>
 [0] Option A<br>
 [1] Option B<br>
 [2] Option C<br>

The correct answer is the function ID (Question no) mod 3, thus 74 (above example)mod 3 = 2, the correct answer is 2. Likewise for ID 75 the answer is 0.

<code>qconsumer qans</code> Which submits the list of answers to the server and renderes the result.

usage : <code>qconsumer qans 74,0 1,2 85,2</code> As can be seen the answers are given as pairs delimetered by a comma, ID,ANS

