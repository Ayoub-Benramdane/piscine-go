const defaultCurry = (obj1)=>(obj2) => {
    let res = {} 
    for (const key in obj1) {
        res[key] = obj1[key]
    }
    for (const key in obj2) {
        res[key] = obj2[key]
    }
    return res
}

const mapCurry = (func)=>(obj) => {
    let res = {}
    for (const key in obj) {
        let [a,b] = func([key,obj[key]]) 
        res[a] = b
    }
    return res
}

const reduceCurry = (func)=>(obj,val=0)=> {
    let k = Object.entries(obj)
    let l = Object.fromEntries(k)
    let res = val
    for (const key in l) {
        res = func(res,[key,l[key]])
    }
    return res
}

const filterCurry = (func) => obj => {
    let res = {}
    for (const key in obj) {
        if (func([key,obj[key]])) {
            res[key] = obj[key]
        }
    }
    return res
}

const reduceScore = (personnel,val) => {
    let ff =(filterCurry((([k,v]) => v.isForceUser))(personnel))
    return reduceCurry((acc, [k,v]) => (acc + v.pilotingScore + v.shootingScore))(ff,val)
}

const filterForce = filterCurry(([k,v]) => v.shootingScore >= 80 && v.isForceUser)

const mapAverage = mapCurry(([k, v]) => {
    const averageScore = (v.pilotingScore + v.shootingScore) / 2
    return [k, {...v, averageScore}]
})