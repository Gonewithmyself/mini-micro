function parse(src){
    var obj = JSON.parse(src)
    , means = obj.dict.word_means;
    return JSON.stringify(means)
}