function showNumberWithAnimation(i, j, num) {
    let $numberCell = $("#number-cell-" + i + "-" + j);
    $numberCell.css("background-color", getNumberBackgroundColor(num));
    $numberCell.css("color", getNumberColor(num));
    $numberCell.text(num)

    $numberCell.animate({
        width: cellSideLength,
        height: cellSideLength,
        top: getPositionTop(i, j),
        left: getPositionLeft(i, j)
    }, 50)
}

function showMoveAnimation(fromX, fromY, toX, toY) {
    let $numberCell = $("#number-cell-" + fromX + "-" + fromY);
    $numberCell.animate({
        top: getPositionTop(toX, toY),
        left: getPositionLeft(toX, toY)
    }, 200);
}

function updateScore(score) {
    $("#score").text(score)
}