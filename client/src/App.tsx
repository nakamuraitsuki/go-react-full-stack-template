import { Route, Routes } from "react-router";
import { Header } from "./components/header";
import { lazy } from "react";
import { Container } from "./components/container/container";
import { AuthGuard } from "./components/auth-guard";

const Top = lazy(() => import("./pages/top"));
const Login = lazy(() => import("./pages/login"));
const SignUp = lazy(() => import("./pages/signup"));
const MyPage = lazy(() => import("./pages/my-page"))

function App() {
  return (
    <>
      <Header />
      <Container>
        <Routes>
          <Route path="/" element={<Top />} />
          <Route
            path="/my-page"
            element={
              <AuthGuard>
                <MyPage />
              </AuthGuard>
            }
          />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
        </Routes>
      </Container>
    </>
  );
}

export default App;
