var mongoose     = require('mongoose');
var Schema       = mongoose.Schema;

var CartSchema   = new Schema({
    cartId : Number,
    groupId : Number,
    productId : Number,
    quantity : Number
});
var cart = mongoose.model('cart', CartSchema);

var ProductSchema   = new Schema({
    productId : Number,
    productName : String,
    description : String,
    price : Number
});
var products = mongoose.model('products', ProductSchema);

module.exports = {
    cart: cart,
    products: products
};



