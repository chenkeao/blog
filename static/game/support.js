let colorDic = {
    2: "#eee4da",
    4: "#ede0c8",
    8: "#f2b179",
    16: "#f59563",
    32: "#f67c5f",
    64: "#f65e3b",
    128: "#edcf72",
    256: "#ebd163",
    512: "#dec14c",
    1024: "#edc53f",
    2048: "#efc32e"
};

let documentWidth = window.screen.availWidth;
let gridContainerWidth = 0.92 * documentWidth;
let cellSideLength = 0.18 * documentWidth;
let cellSpace = 0.04 * documentWidth;

/**
 *
 * @param {number} i 格子的行号
 * @param {number} j 格子的列号
 */
function getPositionTop(i, j) {
    return cellSpace + i * (cellSpace + cellSideLength);
}

/**
 *
 * @param {number} i 格子的行号
 * @param {number} j 格子的列号
 */
function getPositionLeft(i, j) {
    return cellSpace + j * (cellSpace + cellSideLength);
}

/**
 *
 * @param {number} number 当前DIV的值
 */
function getNumberBackgroundColor(number) {
    return colorDic[number]
}

/**
 *
 * @param {number} num 当前DIV的值
 */
function getNumberColor(num) {
    if (num <= 4) {
        return "#776e65";
    }
    return "#fff";
}

/**
 * 判断是否有空格
 */
function noSpace() {
    for (let i = 0; i < 4; i++) {
        for (let j = 0; j < 4; j++) {
            if (board[i][j] === 0) {
                return false;
            }
        }
    }
    return true;
}

/**
 *
 * @param {object} board 判断是否可以向左移动
 */
function canMoveLeft(board) {
    for (let i = 0; i < 4; i++) {
        for (let j = 1; j < 4; j++) {
            if (board[i][j] !== 0) {
                if (board[i][j - 1] === 0 || board[i][j] === board[i][j - 1]) {
                    return true;
                }
            }
        }
    }
    return false;
}

function canMoveRight(board) {
    for (let i = 0; i < 4; i++) {
        for (let j = 2; j >= 0; j--) {
            if (board[i][j] !== 0) {
                if (board[i][j + 1] === 0 || board[i][j] === board[i][j + 1]) {
                    return true;
                }
            }
        }
    }
    return false;
}

function canMoveUp(board) {
    for (let j = 0; j < 4; j++) {
        for (let i = 1; i < 4; i++) {
            if (board[i][j] !== 0) {
                if (board[i - 1][j] === 0 || board[i - 1][j] === board[i][j]) {
                    return true
                }
            }
        }
    }
    return false
}

function canMoveDown(board) {
    for (let j = 0; j < 4; j++) {
        for (let i = 2; i >= 0; i--) {
            if (board[i][j] !== 0) {
                if (board[i + 1][j] === 0 || board[i + 1][j] === board[i][j]) {
                    return true
                }
            }
        }
    }
    return false
}

function noBlockHorizontal(row, col1, col2, board) {
    for (let i = col1 + 1; i < col2; i++) {
        if (board[row][i] !== 0)
            return false;
    }
    return true;
}

function noBlockVertical(coll, row1, row2, board) {
    for (let i = row1 + 1; i < row2; i++) {
        if (board[i][coll] !== 0) {
            return false;
        }
    }
    return true;
}

function copyArray(fromArr, toArr) {
    for (var i = 0; i < fromArr.length; i++) {
        toArr[i] = fromArr[i].concat()
    }
}

let gameHistory = {
    index: 3,
    first: [],
    secend: [],
    third: [],
    update: function (newArr) {
        copyArray(this.secend, this.first);
        copyArray(this.third, this.secend);
        copyArray(newArr, this.third);
        this.index = 3;
    },
    getHistory: function () {
        switch (this.index) {
            case 3:
                return this.third;
            case 2:
                return this.secend;
            case 1:
                return this.first;
            default:
                break;
        }
    },
    clear: function () {
        this.index = 3;
        this.first = [];
        this.secend = [];
        this.third = [];
    }
};
