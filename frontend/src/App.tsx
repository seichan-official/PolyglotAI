import { useEffect, useState } from "react";

function App () {
  const [message,setMessage] = useState('');

  useEffect(()=>{
    //非同期先のURL指定
    fetch("http://localhost:8080/api/hello")
    .then((res) => res.json())
    .then((data) => setMessage(data.message))
    .catch((err) => console.error(err));
  }, []);
  return (
    <div className="p-6 text-center">
      <h1 className="text-2xl font-bold">PolyglotAI</h1>
      <p className="mt-4">Backend says: {message}</p>
    </div>
  );
}
export default App;