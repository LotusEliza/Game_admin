export function  callFunction() {
    let currentDateWithFormat = new Date().toJSON().slice(0, 10).replace(/-/g, '-');
    return currentDateWithFormat
}
