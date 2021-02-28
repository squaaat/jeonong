const isBrowser=new Function("try {return this===window;}catch(e){ return false;}");
// tests if global scope is binded to window

export default isBrowser