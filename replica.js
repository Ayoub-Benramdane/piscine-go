
function fusion(obj1, obj2) {
    const res = {}
    const keys = new Set([...Object.keys(obj1), ...Object.keys(obj2)])

    keys.forEach(key => {
        const vob1 = obj1[key]
        const vob2 = obj2[key]
        // console.log(vob1,vob2)

        if (vob1 !== undefined && vob2 === undefined) {
            res[key] = vob1
        } else if (vob1 === undefined && vob2 !== undefined) {
            res[key] = vob2
        } else {
            

            if (Array.isArray(vob1) && Array.isArray(vob2)) {
                res[key] = vob1.concat(vob2)
            } else if (typeof vob1 === 'string' && typeof vob2 === 'string') {
                res[key] = `${vob1} ${vob2}`
                console.log(res[key])
            } else if (typeof vob1 === 'number' && typeof vob2 === 'number') {
                res[key] = vob2
            } else if (typeof vob1 === 'object' && typeof vob2 === 'object' && !Array.isArray(vob1) && !Array.isArray(vob2)) {
                res[key] = fusion(vob1, vob2)
            } else {
                res[key] = vob2
            }
        }
    })
    return res
}

function replica(to, ...object) {
    if (to === null || typeof to !== 'object') return to
    for (const oneof of object) {
        
        if (oneof !== null && typeof oneof === 'object') {
            for (const key in oneof) {
                if(typeof to === 'object' && typeof oneof === 'object') {
                    to = fusion(to,oneof)
                }
            }
        }
    }
    return to
}


// console.log(replica({ a: 4 }, { a: { b: 1 } }))
// console.log(replica({ a: { b: 1, c: 4 } }, { a: { c: 23 } }))

// console.log(replica(
//     {},
//     Object.freeze({ line: 'Replicants are like any other machine' }),
//     Object.freeze({ author: 'Rich' })))