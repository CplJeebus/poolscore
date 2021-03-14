package utils

const style = `
	<style type="text/css" media="screen"><!--
		* { box-sizing: border-box; }
		body {margin: 0; padding: 0; font-family: Verdana,Geneva,sans-serif;}
		#outer { position: absolute; display: table; width: 100%; height: 99%; margin: 0 auto; }
		#upper { display: table-row; width: 100%; height: 85%; }
		#lower { display: table-row; width: 100%; height: 15%; max-height: 90px; }
		#content { position: relative; display: table-cell; vertical-align: middle; text-align: center; }
		#footer { position: relative; display: table-cell; vertical-align: bottom; text-align: center; }
		.big { font-size: 36pt; }
		--></style>
`
const Pretty = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Eircodes</title>
</head>
<body>
    <div id="myData"></div>
</body>
</html>`

const PrettyForm = `<!DOCTYPE html>
<html>
<body>

<h3>New Code</h3>

<form method="POST" action="/new">
  <label for="fname">Who is it:</label><br>
  <input type="text" id="name" name="name" value=""><br>
  <label for="lname">Code:</label><br>
  <input type="text" id="code" name="code" value=""><br><br>
  <input type="submit" value="Submit">
</form>

</body>
</html>`

const Challenge = `<!DOCTYPE html>
<html>
<head>` + style +
	`</head>
<body>


</body>
</html>`
