function deepCopy(origine) {
    if (origine === null || typeof origine !== 'object' || origine instanceof RegExp) {
        return origine
    }
    if (Array.isArray(origine)) return origine.map(item => deepCopy(item))
    const copiedObject = {}
    for (const key in origine) {
        if (origine.hasOwnProperty(key)) {
            copiedObject[key] = deepCopy(origine[key])
        }
    }
    return copiedObject
}