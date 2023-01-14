"use strict";

var reviews = new Array();

function sendFeedback() {
    const name = document.getElementById("name-input");
    const fBack = document.getElementById("review");

    const nameData = name.value;
    const fBackValue = fBack.value;

    const date = new Date()
    reviews.push([nameData, fBackValue, `${date.getDate()}.${date.getMonth() + 1} ${date.getHours()}:${date.getMinutes()}`]);
    const divFback = document.getElementById("feedback");
    divFback.innerHTML = "";

    for (let item of reviews) {
        const fBackBlock = document.createElement("div");
        fBackBlock.className = "feedback-block";

        const nameDiv = document.createElement("div");
        nameDiv.className = "name-div";
        // 
        nameDiv.innerHTML = `${item[0]} ${item[2]}`;

        const reviewDiv = document.createElement("div");
        reviewDiv.className = "review-value";
        reviewDiv.innerHTML = item[1];

        fBackBlock.appendChild(nameDiv);
        fBackBlock.appendChild(reviewDiv);

        divFback.appendChild(fBackBlock);

    }
}

addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("send-feedback");
    btn.onclick = sendFeedback;
});
