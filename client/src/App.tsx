import { RouterProvider } from "react-router-dom";
import { AppRouter } from "./app-router/app-router";

function App() {
  return (
    <>
      <RouterProvider router={AppRouter} />
    </>
  );
}

export default App;
