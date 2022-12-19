import { Link } from "react-router-dom";
import Product from "./Product";
import ProductH from "./ProductH";
import axios from "axios";
import { useEffect, useMemo, useState, useRef } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import ScrollToTopOnMount from "../template/ScrollToTopOnMount";


function FilterMenuLeft(props) {
  const [cates, setCates] = useState([]);
  const minRef = useRef();
  const maxRef = useRef();
  
  useEffect(() => {
    callAPICategories();
  }, [])

  function callAPICategories() {
    axios.get("http://localhost:3070/api/v1/categories")
    .then((res) => {
      const resp = res.data;
      if (resp.code === "SUCCESS") {
        setCates(resp.data);
      }
    })
    .catch((err) => {
      console.log(err);
    });
  }

  function handleClick(id) {
    props.handleCategory(id);
  }

  function handleApply() {
    const min = minRef.current.value;
    const max = maxRef.current.value;
    props.handlePriceRange(min, max);
  }

  return (
    <ul className="list-group list-group-flush rounded">
      <li className="list-group-item d-none d-lg-block">
        <h5 className="mt-1 mb-2">Categories</h5>
        <div className="d-flex flex-wrap my-2">
          {cates.map((v, i) => {
            return (
              <Link onClick={() => handleClick(v.id)}
                key={v.id}
                to="/products"
                className="btn btn-sm btn-outline-dark rounded-pill me-2 mb-2"
                replace
              >
                {v.name}
              </Link>
            );
          })}
        </div>
      </li>
      <li className="list-group-item">
        <h5 className="mt-1 mb-2">Price Range</h5>
        <div className="d-grid d-block mb-3">
          <div className="form-floating mb-2">
            <input
              type="text"
              className="form-control"
              placeholder="Min"
              defaultValue="100000"
              ref={minRef}
            />
            <label htmlFor="floatingInput">Min Price</label>
          </div>
          <div className="form-floating mb-2">
            <input
              type="text"
              className="form-control"
              placeholder="Max"
              defaultValue="500000"
              ref={maxRef}
            />
            <label htmlFor="floatingInput">Max Price</label>
          </div>
          <button className="btn btn-dark" onClick={handleApply}>Apply</button>
        </div>
      </li>
    </ul>
  );
}

function ProductList() {
  const limit = 6;
  const [viewType, setViewType] = useState({ grid: true });
  const [products, setProducts] = useState([]);
  const [mURL, setURL] = useState("http://localhost:3070/api/v1/products?page=1&limit="+limit+"&order=name:asc");
  const [page, setPage] = useState(1);
  const [order, setOrder] = useState("name:asc");
  const memoProducts = useMemo(() => products, [products]);

  function callAPI(url) {
    axios.get(url)
    .then((res) => {
      const resp = res.data;
      if (resp.code === "SUCCESS") {
        setProducts(resp.data);
        setURL(url)
      }
    })
    .catch((err) => {
      console.log(err);
    });
  }

  function changeViewType() {
    setViewType({
      grid: !viewType.grid,
    });
  }

  function handleChange(event) {
    const order = event.target.value;
    const url = "http://localhost:3070/api/v1/products?page=" + page +"&limit=" +limit+"&order=" + order;
    setOrder(order);
    callAPI(url);
  }

  function handlePageChange(act) {
    let tmpPage = page + act;
    if (tmpPage < 1) tmpPage = 1;
    const url = "http://localhost:3070/api/v1/products?page=" + tmpPage + "&limit="+limit+"&order=" + order;
    callAPI(url);
    setPage(tmpPage);
  }

  function handleSetCategory(category_id) {
    setURL("http://localhost:3070/api/v1/products?page=1&limit="+limit+"&order="+order+"&filter=category_id:eq:" + category_id);
  }

  function handlePriceRange(min, max) {
    setURL("http://localhost:3070/api/v1/products?page=1&limit="+limit+"&order="+order+"&filter=price:gte:" + min + "|price:lte:" + max);
  }

  useEffect(() => {
    callAPI(mURL);
  }, [mURL])


  return (
    <div className="container mt-5 py-4 px-xl-5">
      <ScrollToTopOnMount />
      <nav aria-label="breadcrumb" className="bg-custom-light rounded">
        <ol className="breadcrumb p-3 mb-0">
          <li className="breadcrumb-item">
            <Link
              className="text-decoration-none link-secondary"
              to="/products"
              replace
            >
              All Prodcuts
            </Link>
          </li>
          <li className="breadcrumb-item active" aria-current="page">
            Cases &amp; Covers
          </li>
        </ol>
      </nav>
      <div className="row mb-4 mt-lg-3">
        <div className="d-none d-lg-block col-lg-3">
          <div className="border rounded shadow-sm">
            <FilterMenuLeft handleCategory={handleSetCategory} handlePriceRange={handlePriceRange}/>
          </div>
        </div>
        <div className="col-lg-9">
          <div className="d-flex flex-column h-100">
            <div className="row mb-3">
              <div className="col-lg-3 d-none d-lg-block">
                <select
                  className="form-select"
                  aria-label="Default select example"
                  defaultValue=""
                  onChange={handleChange}
                >
                  <option value="name:asc">Product Name ↑</option>
                  <option value="name:desc">Product Name ↓</option>
                  <option value="price:asc">Price ↑</option>
                  <option value="price:desc">Price ↓</option>
                  <option value="id:asc">Product ID ↑</option>
                  <option value="id:desc">Product ID ↓</option>
                </select>
              </div>
              <div className="col-lg-9 col-xl-5 offset-xl-4 d-flex flex-row">
                <div className="input-group">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="Search products..."
                    aria-label="search input"
                  />
                  <button className="btn btn-outline-dark">
                    <FontAwesomeIcon icon={["fas", "search"]} />
                  </button>
                </div>
                <button
                  className="btn btn-outline-dark ms-2 d-none d-lg-inline"
                  onClick={changeViewType}
                >
                  <FontAwesomeIcon
                    icon={["fas", viewType.grid ? "th-list" : "th-large"]}
                  />
                </button>
              </div>
            </div>
            <div
              className={
                "row row-cols-1 row-cols-md-2 row-cols-lg-2 g-3 mb-4 flex-shrink-0 " +
                (viewType.grid ? "row-cols-xl-3" : "row-cols-xl-2")
              }
            >
              {
                memoProducts.map((product, index) => {
                  if (viewType.grid) {
                    return (
                      <Product key={product.id} percentOff={index % 2 === 0 ? 15 : null} product={product} />
                    );
                  }
                  return (
                    <ProductH key={product.id} percentOff={index % 4 === 0 ? 15 : null} product={product} />
                  );
                })
              }
            </div>
            <div className="d-flex align-items-center mt-auto">
              <span className="text-muted small d-none d-md-inline">
                Showing 10 of 100
              </span>
              <nav aria-label="Page navigation example" className="ms-auto">
                <ul className="pagination my-0">
                  <li className="page-item">
                    <button className="page-link" onClick={() => handlePageChange(-1)}>
                      Previous
                    </button>
                  </li>
                  <li className="page-item">
                    <button className="page-link" onClick={() => handlePageChange(1)}>
                      Next
                    </button>
                  </li>
                </ul>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ProductList;
