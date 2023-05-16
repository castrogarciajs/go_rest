function App() {
  const GET = async () => {
    const res = await fetch("http://localhost:3000/users");

    const data = res.json();

    console.log(data);
  };
  return (
    <div>
      <button onClick={GET}>JSON</button>
    </div>
  );
}

export default App;
