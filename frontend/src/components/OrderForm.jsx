import { useState } from 'react'

function OrderForm({ prices, onSubmit }) {
  const [symbol, setSymbol] = useState('')
  const [side, setSide] = useState('buy')
  const [quantity, setQuantity] = useState('')
  const [price, setPrice] = useState('')

  const handleSubmit = (e) => {
    e.preventDefault()

    if (!symbol || !quantity || !price) {
      alert('Please fill in all fields')
      return
    }

    const order = {
      symbol,
      side,
      quantity: parseInt(quantity),
      price: parseFloat(price),
    }

    if (onSubmit(order)) {
      // Reset form
      setSymbol('')
      setQuantity('')
      setPrice('')
    }
  }

  const handleSymbolChange = (e) => {
    const selectedSymbol = e.target.value
    setSymbol(selectedSymbol)
    
    // Auto-fill current price if available
    if (selectedSymbol && prices[selectedSymbol]) {
      setPrice(prices[selectedSymbol].price.toFixed(2))
    } else {
      setPrice('')
    }
  }

  return (
    <div className="bg-white rounded-xl p-6 shadow-md">
      <h2 className="mb-5 text-gray-800 text-2xl font-semibold">Place Order</h2>
      <form onSubmit={handleSubmit} className="space-y-5">
        <div>
          <label htmlFor="symbol" className="block mb-2 font-semibold text-gray-600">Symbol</label>
          <select
            id="symbol"
            value={symbol}
            onChange={handleSymbolChange}
            required
            className="w-full px-3 py-3 border-2 border-gray-200 rounded-lg text-base focus:outline-none focus:border-indigo-500 transition-colors"
          >
            <option value="">Select a stock</option>
            {Object.keys(prices).map((sym) => (
              <option key={sym} value={sym}>
                {sym}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label htmlFor="side" className="block mb-2 font-semibold text-gray-600">Side</label>
          <div className="flex gap-3">
            <button
              type="button"
              className={`flex-1 px-3 py-3 border-2 rounded-lg text-base font-semibold transition-all ${
                side === 'buy'
                  ? 'bg-green-500 text-white border-green-500'
                  : 'bg-white text-gray-700 border-gray-200 hover:border-indigo-500'
              }`}
              onClick={() => setSide('buy')}
            >
              Buy
            </button>
            <button
              type="button"
              className={`flex-1 px-3 py-3 border-2 rounded-lg text-base font-semibold transition-all ${
                side === 'sell'
                  ? 'bg-red-500 text-white border-red-500'
                  : 'bg-white text-gray-700 border-gray-200 hover:border-indigo-500'
              }`}
              onClick={() => setSide('sell')}
            >
              Sell
            </button>
          </div>
        </div>

        <div>
          <label htmlFor="quantity" className="block mb-2 font-semibold text-gray-600">Quantity</label>
          <input
            type="number"
            id="quantity"
            value={quantity}
            onChange={(e) => setQuantity(e.target.value)}
            min="1"
            required
            className="w-full px-3 py-3 border-2 border-gray-200 rounded-lg text-base focus:outline-none focus:border-indigo-500 transition-colors"
          />
        </div>

        <div>
          <label htmlFor="price" className="block mb-2 font-semibold text-gray-600">Price</label>
          <input
            type="number"
            id="price"
            value={price}
            onChange={(e) => setPrice(e.target.value)}
            step="0.01"
            min="0.01"
            required
            className="w-full px-3 py-3 border-2 border-gray-200 rounded-lg text-base focus:outline-none focus:border-indigo-500 transition-colors"
          />
          {symbol && prices[symbol] && (
            <span className="block mt-1.5 text-sm text-gray-500">
              Current: ${prices[symbol].price.toFixed(2)}
            </span>
          )}
        </div>

        <button
          type="submit"
          className="w-full px-4 py-3.5 bg-gradient-to-r from-indigo-500 to-purple-600 text-white rounded-lg text-lg font-semibold hover:shadow-lg hover:-translate-y-0.5 active:translate-y-0 transition-all"
        >
          Place {side === 'buy' ? 'Buy' : 'Sell'} Order
        </button>
      </form>
    </div>
  )
}

export default OrderForm

