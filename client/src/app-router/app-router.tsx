import { createBrowserRouter } from "react-router-dom";
import { homePage } from "./routes/home-page/home-page";
import { anyRoute } from "./routes/any-route/any-route";
import { notFound } from "./routes/not-found/not-found";

export const AppRouter = createBrowserRouter([homePage, anyRoute, notFound]);
