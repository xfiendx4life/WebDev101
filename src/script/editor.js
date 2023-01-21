"use strict";

addEventListener("DOMContentLoaded", () => {
    const container = document.getElementsByClassName("container")[0];
    container.classList.add("editor-container");

    const tools = document.createElement("div");
    tools.className = "tools";
    tools.id = "tools";
    container.appendChild(tools);

    const boldBtn = document.createElement("button");
    boldBtn.className = "btn btn-primary";
    boldBtn.innerHTML = "bold";
    tools.appendChild(boldBtn);

    const textAria = document.createElement("textarea");
    textAria.className = "form-control";
    textAria.classList.add("editor-space");
    container.appendChild(textAria);

    const res = document.createElement("div");
    res.className = "result";
    res.id = "result";
    container.appendChild(res);

    textAria.oninput = () => {
        res.innerHTML = textAria.value;
    }

    textAria.onchange = () => {
        res.innerHTML = textAria.value;
    }

    boldBtn.onclick = () => {
        const selStart = textAria.selectionStart;
        const selEnd = textAria.selectionEnd;
        let boldSubs = `${textAria.value.substring(0, selStart)}<b>${textAria.value.substring(selStart, selEnd)}</b>`
        textAria.value = boldSubs;
    }; 
    const aBtn = new aButton(textAria);
    tools.appendChild(aBtn.container)

});

 /**
  * creates HyperLink button
  * @param {HTMLTextAreaElement} textarea 
  */
function aButton(textAria) {
    this.container = document.createElement("div");
    this.container.className = "hyperlink-container";
    const button = document.createElement('button');
    button.className = 'btn btn-primary';
    button.classList.add("a-button");
    button.id = 'a-button';
    button.innerHTML = 'Hyperlink';
    this.container.appendChild(button);

    const hrefInput = document.createElement('input');
    hrefInput.setAttribute('type', "text");
    hrefInput.id = "href-input";
    hrefInput.className = 'form-control';
    hrefInput.style.display = 'none';
    hrefInput.placeholder = "Write href attribute";
    this.container.appendChild(hrefInput);


    hrefInput.onchange = () => {
        const href = hrefInput.value;
        const selStart = textAria.selectionStart;
        const selEnd = textAria.selectionEnd;
        let boldSubs = `${textAria.value.substring(0, selStart)}<a href="${href}">${textAria.value.substring(selStart, selEnd)}</a>`
        textAria.value = boldSubs;
        hrefInput.style.display = "none";
    }

    button.onclick = () => {
        hrefInput.style.display = "block";
    }
}
