import styles from "../styles/Home.module.css";
import { useState } from "react";
import axios from "axios";
export default function Home() {
  const [ProductId, setProductId] = useState("");
  const [StoreId, setStoreId] = useState("");
  const [Title, setProductName] = useState("");
  const [Tag, setTag] = useState("");
  const [Price, setPrice] = useState("");
  const [Quantity, setQuantity] = useState("");
  const [LocationN, setLocation] = useState("");
  const postData = () => {
    const jsonData = {
      ProductId,
      StoreId,
      Title,
      Tag,
      Price,
      Quantity,
      LocationN,
    };
    axios.post("http://localhost:8080/", jsonData);
  };
  return (
    <div
      className={
        "p-6 max-w-sm mx-auto bg-white rounded-xl shadow-md space-x-4 " +
        styles.color
      }
    >
      <input
        value={ProductId}
        placeholder="productId"
        onChange={(e) => setProductId(e.target.value)}
      />
      <input
        value={StoreId}
        placeholder="StoreId"
        onChange={(e) => setStoreId(e.target.value)}
      />
      <input
        value={Title}
        placeholder="Product Name"
        onChange={(e) => setProductName(e.target.value)}
      />
      <input
        value={Tag}
        placeholder="tag"
        onChange={(e) => setTag(e.target.value)}
      />
      <input
        value={Price}
        placeholder="price"
        onChange={(e) => setPrice(e.target.value)}
      />
      <input
        value={Quantity}
        placeholder="Quantity"
        onChange={(e) => setQuantity(e.target.value)}
      />
      <input
        value={LocationN}
        placeholder="Location"
        onChange={(e) => setLocation(e.target.value)}
      />
      <button
        className="py-2 px-4 font-semibold rounded-lg shadow-md text-white bg-green-500 hover:bg-green-700"
        onClick={postData}
      >
        Click me
      </button>
    </div>
  );
}
