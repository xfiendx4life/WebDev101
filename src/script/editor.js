"use strict";

addEventListener("DOMContentLoaded", () => {
    const container = document.getElementsByClassName("container")[0];

    const tools = document.createElement("div");
    tools.className = "tools";
    tools.id = "tools";
    container.appendChild(tools);

    const boldBtn = document.createElement("button");
    boldBtn.className = "btn btn-primary";
    boldBtn.innerHTML = "bold";
    boldBtn.onclick = () => {}; // TODO: написать обработчик, для создания жирного текста
    tools.appendChild(boldBtn);

    const textAria = document.createElement("textarea");
    textAria.className = "form-control";
    textAria.classList.add("editor-space");
    container.appendChild(textAria);
});
