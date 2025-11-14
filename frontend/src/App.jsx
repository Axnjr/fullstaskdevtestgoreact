import { useState, useEffect } from 'react'
import StockPrices from './components/StockPrices'
import OrderForm from './components/OrderForm'
import OrdersList from './components/OrdersList'
import Login from './components/Login'

const API_BASE = 'http://localhost:8080'

function App() {
  const [prices, setPrices] = useState({})
  const [orders, setOrders] = useState([])
  const [ws, setWs] = useState(null)
  const [token, setToken] = useState(null)
  const [user, setUser] = useState(null)

  // check for existing token on mount
  useEffect(() => {
    const savedToken = localStorage.getItem('token')
    const savedUser = localStorage.getItem('user')
    if (savedToken && savedUser) {
      setToken(savedToken)
      setUser(JSON.parse(savedUser))
    }
  }, [])

  useEffect(() => {
    if (token) {
      fetchPrices()
      fetchOrders()
    }
  }, [token])

  // ws only when authenticated
  useEffect(() => {
    if (!token) {
      // close any existing ws if token is removed
      if (ws) {
        ws.close()
        setWs(null)
      }
      return
    }

    const websocket = new WebSocket('ws://localhost:8080/ws')
    
    websocket.onopen = () => {
      console.log('WebSocket connected')
    }

    websocket.onmessage = (event) => {
      try {
        const update = JSON.parse(event.data)
        setPrices(prev => ({
          ...prev,
          [update.symbol]: {
            symbol: update.symbol,
            price: update.price,
            change: update.change,
            changePct: update.changePct
          }
        }))
      } catch (error) {
        console.error('Error parsing WebSocket message:', error)
      }
    }

    websocket.onerror = (error) => {
      console.error('WebSocket error:', error)
    }

    websocket.onclose = () => {
      console.log('WebSocket disconnected')
    }

    setWs(websocket)

    return () => {
      if (websocket.readyState === WebSocket.OPEN || websocket.readyState === WebSocket.CONNECTING) {
        websocket.close()
      }
    }
  }, [token])

  const handleLogin = (newToken, newUser) => {
    setToken(newToken)
    setUser(newUser)
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    setToken(null)
    setUser(null)
    setOrders([])
    if (ws) {
      ws.close()
      setWs(null)
    }
  }

  const fetchPrices = async () => {
    try {
      const response = await fetch(`${API_BASE}/prices`)
      const data = await response.json()
      const pricesMap = {}
      data.forEach(stock => {
        pricesMap[stock.symbol] = {
          symbol: stock.symbol,
          price: stock.price,
          change: 0,
          changePct: 0
        }
      })
      setPrices(pricesMap)
    } catch (error) {
      console.error('Error fetching prices:', error)
    }
  }

  const fetchOrders = async () => {
    if (!token) {
      setOrders([])
      return
    }
    
    try {
      const response = await fetch(`${API_BASE}/orders`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (response.status === 401) {
        handleLogout()
        setOrders([])
        return
      }

      if (!response.ok) {
        console.error('Failed to fetch orders:', response.status)
        setOrders([])
        return
      }

      const data = await response.json()
      setOrders(data ?? [])
    } catch (error) {
      console.error('Error fetching orders:', error)
      setOrders([])
    }
  }

  const handleOrderSubmit = async (order) => {
    if (!token) {
      alert('Please login to place orders')
      return false
    }

    try {
      const response = await fetch(`${API_BASE}/orders`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(order),
      })

      if (response.status === 401) {
        handleLogout()
        alert('Session expired. Please login again.')
        return false
      }

      if (response.ok) {
        const newOrder = await response.json()
        setOrders(prev => [newOrder, ...prev])
        return true
      } else {
        const error = await response.json()
        alert(error.error || 'Failed to place order')
        return false
      }
    } catch (error) {
      console.error('Error placing order:', error)
      alert('Failed to place order')
      return false
    }
  }

  // show login if not authenticated
  if (!token) {
    return <Login onLogin={handleLogin} />
  }

  return (
    <div className="min-h-screen p-5">
      <header className="text-center text-white mb-8">
        <div className="flex justify-between items-center max-w-7xl mx-auto mb-4">
          <div></div>
          <div>
            <h1 className="text-4xl font-bold mb-2 drop-shadow-lg">📈 Trading Dashboard</h1>
            <p className="text-lg opacity-90">Real-time stock prices and order management</p>
          </div>
          <div className="flex items-center gap-4">
            <div className="text-right">
              <p className="text-sm opacity-80">Logged in as</p>
              <p className="font-semibold">{user?.username}</p>
            </div>
            <button
              onClick={handleLogout}
              className="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg font-semibold transition-colors backdrop-blur-sm"
            >
              Logout
            </button>
          </div>
        </div>
      </header>

      <div className="grid grid-cols-1 lg:grid-cols-[1fr_400px] gap-5 max-w-7xl mx-auto">
        <div className="flex flex-col gap-5">
          <StockPrices prices={prices} />
          <OrderForm 
            prices={prices} 
            onSubmit={handleOrderSubmit}
          />
        </div>
        
        <div className="flex flex-col">
          <OrdersList orders={orders} />
        </div>
      </div>
    </div>
  )
}

export default App

