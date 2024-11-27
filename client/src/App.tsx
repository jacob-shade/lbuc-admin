import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import Home from "./pages/Home"
import Team from "./pages/Team"
import { Layout } from "@/components/layout"

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/team/:id" element={<Team />} />
          <Route path="*" element={<div>404 Not Found</div>} />
        </Routes>
      </Layout>
    </Router>
  )
}

export default App
