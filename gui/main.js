var text_to_type;
var cur_idx;
var code_with_spaces = false;
var hits;
var fails;
var let_errors = false;
var startTime;
var endTime = "lol";

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
}

// handle tabs in textarea
$("textarea").keydown(function(e) {
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
			if (text_to_type.length - 1 <= cur_idx) {
				document.getElementById("score").style.color="#9f6";
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
	document.getElementById("score").innerHTML="score: "+cur_idx+"\\"+(text_to_type.length-1)+";\terrors: "+fails +";\terror rate: " +(fails/hits).toFixed(2)+suffix;
});
