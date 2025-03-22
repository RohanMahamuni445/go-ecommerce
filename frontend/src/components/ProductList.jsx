import React from 'react';

const ProductList = ({ products }) => {
  return (
    <div className="content">
      <h2>Product List</h2>
      <div className="product-list">
        {products.map((product, index) => (
          <div className="product-card" key={index}>
            <h3>{product.name}</h3>
            <p>{product.description}</p>
            <p className="price">Price: ${product.price}</p>
            <p className="stock">Stock: {product.stock}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProductList;

