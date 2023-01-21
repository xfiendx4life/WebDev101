"use strict";

const hyperlink = 'hyperlink';
const image = 'image';

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

    const imgButton = new aButton(textAria, image);
    tools.appendChild(imgButton.container);

});

 /**
  * creates HyperLink button
  * @param {HTMLTextAreaElement} textarea 
  * @param {String} type type of the button to create
  */
function aButton(textAria, type=hyperlink) {
    this.container = document.createElement("div");
    this.container.className = "hyperlink-container";
    const button = document.createElement('button');
    button.className = 'btn btn-primary';
    button.classList.add("a-button");
    button.id = 'a-button';
    let name = 'Hyperlink';
    let hrefOrSrc = 'href';
    if (type == image) {
        name = 'Image';
        hrefOrSrc = 'src'
    }
    button.innerHTML = name;
    this.container.appendChild(button);

    const hrefInput = document.createElement('input');
    hrefInput.setAttribute('type', "text");
    hrefInput.id = "href-input";
    hrefInput.className = 'form-control';
    hrefInput.style.display = 'none';
    hrefInput.placeholder = `Write ${hrefOrSrc} attribute`;
    this.container.appendChild(hrefInput);


    hrefInput.onchange = () => {
        let tagName = 'a';
        const href = hrefInput.value;
        const selStart = textAria.selectionStart;
        const selEnd = textAria.selectionEnd;
        let endOfHyperlink = `${textAria.value.substring(selStart, selEnd)}</${tagName}>`
        if (type == image) {
            tagName = 'img'
            endOfHyperlink = '';
        }
        let boldSubs = `${textAria.value.substring(0, selStart)}<${tagName} ${hrefOrSrc}="${href}">${endOfHyperlink}`
        textAria.value = boldSubs;
        hrefInput.style.display = "none";
    }

    button.onclick = () => {
        if (hrefInput.style.display == 'none') {
            hrefInput.style.display = "block";
        } else {
            hrefInput.style.display = "none";
        }
        
    }
}
