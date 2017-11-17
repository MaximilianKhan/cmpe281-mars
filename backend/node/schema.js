var mongoose     = require('mongoose');
var Schema       = mongoose.Schema;

var CartSchema   = new Schema({
    cartId : String,
    groupId : String,
    productId : String,
    quantity : String
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



