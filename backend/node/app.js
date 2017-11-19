var express = require('express');
var bodyParser = require('body-parser');
var cors = require('cors');
var app = express();
var mongoose = require('mongoose');
var schemas = require('./schema');
var cart = schemas.cart;
var products = schemas.products;
var port = 3000;
var router = express.Router();

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json({ type: 'application/vnd.api+json' }))

// parse application/json

mongoose.connect('mongodb://localhost:27017/cmpe281project');

router.use(function (req, res, next) {
  // do logging 
  console.log('Logging of request will be done here');
  next(); 
});

router.route('/cart').get(function (req, res) {
  console.log('GET cart');
  cart.find(function (err, carts) {
    if (err) {
        res.send(err);
    }
    res.send(carts);
  });
});

router.route('/cart').post(function (req, res) {
  console.log("in add");
  console.log("in add:"+req);
  console.log("in add:"+req.body.cartId);
  
  var p = new cart();
  p.cartId = req.body.cartId;
  p.groupId = req.body.groupId;
  p.productId = req.body.productId;
  p.quantity = req.body.quantity;
  
  p.save(function (err) {
      if (err) {
          res.send(err);
      }
      console.log("added");
      res.send({ message: 'Cart Created !' })
  })
});

router.route('/cart').post(function (req, res) {
  console.log("in add");
  console.log("in add:"+req);
  console.log("in add:"+req.body.cartId);
  
  var p = new cart();
  p.cartId = req.body.cartId;
  p.groupId = req.body.groupId;
  p.productId = req.body.productId;
  p.quantity = req.body.quantity;
  
  p.save(function (err) {
      if (err) {
          res.send(err);
      }
      console.log("added");
      res.send({ message: 'Cart Created !' })
  })
});


router.route('/products').get(function (req, res) {
  console.log('GET products');
  products.find(function (err, products) {
      if (err) {
          res.send(err);
      }
      res.send(products);
  });
});

router.route('/products/:productId').get(function (req, res) {
    products.findById(req.params.productId, function (err, products) {
        if (err)
            res.send(err);
        res.json(products);
    });
});


app.use(cors());
app.use('/api', router);
app.listen(port, () =>{
  console.log(`Server started on port ${port}`);
});
