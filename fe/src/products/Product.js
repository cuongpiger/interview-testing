import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function Product(props) {
  const price = 10000;
  let percentOff;
  let offPrice = `${price}Ks`;

  if (props.percentOff && props.percentOff > 0) {
    percentOff = (
      <div
        className="badge bg-dim py-2 text-white position-absolute"
        style={{ top: "0.5rem", right: "0.5rem" }}
      >
        {props.percentOff}% OFF
      </div>
    );

    offPrice = (
      <>
        <del>{price}Ks</del> {price - (props.percentOff * price) / 100}Ks
      </>
    );
  }

  return (
    <div className="col">
      <div className="card shadow-sm">
        <Link to="/products/1" href="!#" replace>
          <img
            className="card-img-top bg-dark cover"
            height="200"
            alt=""
            src={props.product.images[0]}
          />
        </Link>
        <div className="card-body">
          <h5 className="card-title text-center text-dark text-truncate">
            {props.product.name}
          </h5>
        <p className="card-text text-center text-muted mb-0">{props.product.price} <b>₫</b></p>
          <div className="d-grid d-block">
            <button className="btn btn-outline-dark mt-3">
              <FontAwesomeIcon icon={["fas", "cart-plus"]} /> Add to cart
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Product;
