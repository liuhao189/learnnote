module.exports = (options) => {
    return (req, res, next) => {
        req.options = options;
        next();
    }
}