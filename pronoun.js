let pronouns = ['i','you','he','she','it','they','we']

const words = (arr,ind) => {
        if (!pronouns.includes(arr[ind]))
            return arr[ind]
        return ''
}
function splitString(input) {
    return input.split(/[\s,]+/).map(word => word.trim()).filter(word => word !== '');
}

function pronoun(string){
    let arr = splitString(string)
    // console.log(arr)
    let rtn = {}
    for (let i = 0;i< arr.length;i++){
        if (pronouns.includes(arr[i].toLowerCase())){
            let wwor = words(arr,i+1)
            if (!rtn[arr[i].toLowerCase()]) {
                rtn[arr[i].toLowerCase()] = {word:[],count:1}
                if (wwor != '' && wwor != undefined) rtn[arr[i].toLowerCase()].word.push(wwor)
            }
            else {
                if (wwor != '') rtn[arr[i].toLowerCase()].word.push(wwor)
                rtn[arr[i].toLowerCase()].count++
            }
        }
    }
    return rtn
}



// console.log(pronoun('I buy,\ni to,\nYOU buy,\nit have,\nIt buys,\nit is,\nyou go'))
// console.log(pronoun('we will rock you'))