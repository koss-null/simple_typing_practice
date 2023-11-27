let textToType;
let curIdx;
let codeWithSpaces = !(document.getElementById('use_tabs').checked);
let hits;
let maxSymbols;
let fails;
let letErrors = false;
let startTime;
let endTime = "lol";

function changeUseTabs() {
  codeWithSpaces = !(document.getElementById('use_tabs').checked);
  if (!codeWithSpaces) {
    document.getElementsByClassName('use_tabs')[0].style.backgroundColor = "#dbd851";
    document.getElementsByClassName('use_tabs')[0].style.color = "#333";
    document.getElementsByName('label_use_tabs')[0].innerHTML = "Using tabs instead of spaces";
    textToType = textToType.replace(/    /g, "\t");
    restart();
  } else {
    document.getElementsByClassName('use_tabs')[0].style.backgroundColor = "#444";
    document.getElementsByClassName('use_tabs')[0].style.color = "#dbd851";
    document.getElementsByName('label_use_tabs')[0].innerHTML = "Using spaces instead of tabs";
    textToType = textToType.replace(/\t/g, "    ");
    restart();
  }
}

function changeMaxSymbols() {
  const msOld = maxSymbols;
  maxSymbols = document.getElementById("max_symbols").value;
  if (maxSymbols <= 0) {
    maxSymbols = textToType.length - 1;
  }
  if (msOld !== maxSymbols) {
    restart();
  }
}

function changeCode() {
  const fileList = myFile.files;
  const file = fileList[0];
  const reader = new FileReader();

  reader.addEventListener("load", () => {
    document.getElementsByName('main_window_input')[0].value = "";
    curIdx = 0;
    hits = 0;
    fails = 0;
    if (codeWithSpaces) {
      document.getElementsByName('main_window_text')[0].value = reader.result.replace(/\t/g, "    ");
      textToType = reader.result.replace(/\t/g, "    ");
    } else {
      document.getElementsByName('main_window_text')[0].value = reader.result.replace(/    /g, "\t");
      textToType = reader.result.replace(/    /g, "\t");
    }

    document.getElementById("max_symbols").value = textToType.length - 1;

    const linesCnt = reader.result.split(/\r\n|\r|\n/).length;
    document.getElementsByName('main_window_text')[0].style.minHeight = linesCnt * 27 + "px";
    document.getElementsByName('main_window_input')[0].style.minHeight = linesCnt * 27 + "px";
    document.getElementsByName('main_window_input')[0].style.marginTop = -linesCnt * 27 + "px";
    console.log("min-height set to ", linesCnt * 27, "px");
  }, false);

  if (file) {
    reader.readAsText(file);
  }
}

function restart() {
  hits = 0;
  fails = 0;
  endTime = 'lol';
  curIdx = 0;
  document.getElementsByName('main_window_input')[0].value = "";
  document.getElementsByClassName('content-front')[0].style.borderBlockColor = '#99974a';
  document.getElementsByClassName('content-front')[0].style.background = '#333';
}

// handle tabs in textarea
$("textarea").keydown(function(e) {
  changeMaxSymbols();

  if (hits === 0) {
    startTime = performance.now();
  }
  endTime = performance.now();
  document.getElementById("score").style.color = "#dbd851";
  if (e.code === "Backspace") {
    curIdx -= 1;
    hits -= 1;
    if (curIdx < 0) {
      curIdx = 0;
      startTime = performance.now();
      endTime = "lol";
    }
    if (hits < 0) {
      hits = 0;
    }
    console.log("backspace; hits: ", hits);
  } else {
    const typed = e.key;
    console.log("cur idx ", curIdx, " val: ", typed);
    if (typed === textToType[curIdx]) {
      console.log("got ", typed, " AS expected ", textToType[curIdx]);
      hits += 1;
      curIdx += 1;
    } else if (typed === "Shift") {
      return;
    } else if (typed === "Enter" && textToType[curIdx] === "\n") {
      if (maxSymbols <= curIdx) {
        document.getElementById("score").style.color = "#9f6";
        document.getElementsByClassName('content-front')[0].style.borderBlockColor = '#49904e';
        document.getElementsByClassName('content-front')[0].style.background = '#284c32';
        const start = this.selectionStart;
        const end = this.selectionEnd;

        const $this = $(this);
        const value = $this.val();

        // set textarea value to: text before caret + tab + text after caret
        $this.val(value.substring(0, start) + value.substring(end));

        // put caret at right position again (add one for the tab)
        this.selectionStart = this.selectionEnd = start;

        // prevent the focus lose
        e.preventDefault();
      } else {
        hits += 1;
        curIdx += 1;
      }
    } else if (typed !== "Tab") {
      fails += 1;
      console.log("got ", typed, " expected ", textToType[curIdx]);
      // do not enter current symbol
      const start = this.selectionStart;
      const end = this.selectionEnd;

      const $this = $(this);
      const value = $this.val();

      // set textarea value to: text before caret + tab + text after caret
      $this.val(value.substring(0, start) + value.substring(end));

      // put caret at right position again (add one for the tab)
      this.selectionStart = this.selectionEnd = start;

      // prevent the focus lose
      e.preventDefault();
    }
  }
  console.log(curIdx, " hits ", hits, " fails ", fails);

  if (e.keyCode === 9) { // tab was pressed
    let toAdd = "";
    if (textToType[curIdx] === "\t") {
      curIdx += 1;
      toAdd = "\t";
    }
    // get caret position/selection
    const start = this.selectionStart;
    const end = this.selectionEnd;

    const $this = $(this);
    const value = $this.val();

    // set textarea value to: text before caret + tab + text after caret
    $this.val(value.substring(0, start) + toAdd + value.substring(end));

    // put caret at right position again (add one for the tab)
    this.selectionStart = this.selectionEnd = start + 1;

    // prevent the focus lose
    e.preventDefault();
  }

  let suffix = "";
  if (endTime !== "lol") {
    suffix = `; total time: ${((endTime - startTime) / 1000).toFixed(2)}s,\t${(curIdx / (endTime - startTime) * 1000 * 60).toFixed(2)} char/min`;
  }
  document.getElementById("score").style.marginTop = "15px";
  document.getElementById("score").innerHTML = "score: " + curIdx + "/" + maxSymbols + ";\terrors: " + fails + ";\terror rate: " + (fails / hits).toFixed(2) + suffix;
});
