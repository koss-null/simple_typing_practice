var text_to_type;
var cur_idx;
var code_with_spaces = !(document.getElementById('use_tabs').checked);
var hits;
var max_symbols;
var fails;
var let_errors = false;
var startTime;
var endTime = "lol";


function changeUseTabs() {
	code_with_spaces = !(document.getElementById('use_tabs').checked);
	if (!code_with_spaces) {
		document.getElementsByClassName('use_tabs')[0].style.backgroundColor="#dbd851";
		document.getElementsByClassName('use_tabs')[0].style.color="#333";
		document.getElementsByName('label_use_tabs')[0].innerHTML="Using tabs instead of spaces";
		text_to_type = text_to_type.replace(/    /g, "\t");
		restart();
	} else {
		document.getElementsByClassName('use_tabs')[0].style.backgroundColor="#444";
		document.getElementsByClassName('use_tabs')[0].style.color="#dbd851";
		document.getElementsByName('label_use_tabs')[0].innerHTML="Using spaces instead of tabs";
		text_to_type = text_to_type.replace(/\t/g, "    ");
		restart();
	}
}

function changeMaxSymbols() {
	ms_old = max_symbols
	max_symbols = document.getElementById("max_symbols").value;
	if (max_symbols <= 0) {
		max_symbols = text_to_type.length - 1;
	}
	if (ms_old != max_symbols) {
		restart()
	}
}

function changeCode() {
	var fileList = myFile.files;
	const file = fileList[0];
	let reader = new FileReader();

	reader.addEventListener("load", () => {
		document.getElementsByName('main_window_input')[0].value="";
		cur_idx = 0;
		hits = 0;
		fails = 0;
		if (code_with_spaces) {
			document.getElementsByName('main_window_text')[0].value=reader.result.replace(/\t/g, "    ");
			text_to_type = reader.result.replace(/\t/g, "    ");
		} else {
			document.getElementsByName('main_window_text')[0].value=reader.result.replace(/    /g, "\t");
			text_to_type = reader.result.replace(/    /g, "\t");
		}

		document.getElementById("max_symbols").value = text_to_type.length - 1;

		let linesCnt = reader.result.split(/\r\n|\r|\n/).length;
		document.getElementsByName('main_window_text')[0].style.minHeight=linesCnt*27+"px";
		document.getElementsByName('main_window_input')[0].style.minHeight=linesCnt*27+"px";
		document.getElementsByName('main_window_input')[0].style.marginTop=-linesCnt*27+"px";
		console.log("min-height set to ", linesCnt*27, "px");
	}, false);

	if (file) {
		reader.readAsText(file);
	}
}

function restart() {
	hits = 0;
	fails = 0;
	endTime = 'lol';
	cur_idx = 0;
	document.getElementsByName('main_window_input')[0].value="";
	document.getElementsByClassName('content-front')[0].style.borderBlockColor = '#99974a';
	document.getElementsByClassName('content-front')[0].style.background = '#333';
}

// handle tabs in textarea
$("textarea").keydown(function(e) {
	changeMaxSymbols()

	if (hits == 0) {
		startTime = performance.now();
	}
	endTime = performance.now();
	document.getElementById("score").style.color="#dbd851";
	if (e.code === "Backspace") {
		cur_idx -= 1;
		hits -= 1;
		if (cur_idx < 0) {
			cur_idx = 0;
			startTime = performance.now();
			endTime = "lol";
		}
		if (hits < 0) {
			hits = 0;
		}
		console.log("backspace; hits: ", hits);
	} else {
		var typed = e.key;
		console.log("cur idx ", cur_idx, " val: ", typed);
		if (typed === text_to_type[cur_idx]) {
			console.log("got ", typed, " AS expected ", text_to_type[cur_idx]);
			hits += 1;
			cur_idx += 1;
		} else if (typed === "Shift") {
			return
		} else if (typed === "Enter" && text_to_type[cur_idx] == "\n") {
			if (max_symbols <= cur_idx) {
				document.getElementById("score").style.color="#9f6";
				document.getElementsByClassName('content-front')[0].style.borderBlockColor = '#49904e';
				document.getElementsByClassName('content-front')[0].style.background = '#284c32';
				var start = this.selectionStart;
				var end = this.selectionEnd;

				var $this = $(this);
				var value = $this.val();

				// set textarea value to: text before caret + tab + text after caret
				$this.val(value.substring(0, start)
				    + value.substring(end));

				// put caret at right position again (add one for the tab)
				this.selectionStart = this.selectionEnd = start;

				// prevent the focus lose
				e.preventDefault();
			} else {
				hits += 1;
				cur_idx += 1;
			}
		} else if (typed != "Tab") {
			fails += 1;
			console.log("got ", typed, " expected ", text_to_type[cur_idx]);
			// do not enter current symbol
			var start = this.selectionStart;
			var end = this.selectionEnd;

			var $this = $(this);
			var value = $this.val();

			// set textarea value to: text before caret + tab + text after caret
			$this.val(value.substring(0, start)
			    + value.substring(end));

			// put caret at right position again (add one for the tab)
			this.selectionStart = this.selectionEnd = start;

			// prevent the focus lose
			e.preventDefault();
		}
	}
	console.log(cur_idx, " hits ", hits, " fails ", fails);

	if(e.keyCode === 9) { // tab was pressed
		var to_add = "";
		if (text_to_type[cur_idx] === "\t") {
			cur_idx += 1;
			to_add = "\t";
		}
		// get caret position/selection
		var start = this.selectionStart;
		var end = this.selectionEnd;

		var $this = $(this);
		var value = $this.val();

		// set textarea value to: text before caret + tab + text after caret
		$this.val(value.substring(0, start)
			    + to_add
			    + value.substring(end));

		// put caret at right position again (add one for the tab)
		this.selectionStart = this.selectionEnd = start + 1;

		// prevent the focus lose
		e.preventDefault();
	}

	var suffix = "";
	if (endTime != "lol") {
		suffix = `; total time: ${((endTime - startTime)/1000).toFixed(2)}s,\t${(cur_idx / (endTime - startTime) * 1000 * 60).toFixed(2)} char/min`;
	}
	document.getElementById("score").innerHTML="score: "+cur_idx+"/"+(max_symbols)+";\terrors: "+fails +";\terror rate: " +(fails/hits).toFixed(2)+suffix;
});
