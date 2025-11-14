function StockPrices({ prices }) {
  const stocks = Object.values(prices).sort((a, b) => a.symbol.localeCompare(b.symbol))

  const formatPrice = (price) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2,
      maximumFractionDigits: 2,
    }).format(price)
  }

  const formatChange = (change) => {
    const sign = change >= 0 ? '+' : ''
    return `${sign}${change.toFixed(2)}`
  }

  const formatChangePct = (changePct) => {
    const sign = changePct >= 0 ? '+' : ''
    return `${sign}${changePct.toFixed(2)}%`
  }

  return (
    <div className="bg-white rounded-xl p-6 shadow-md">
      <h2 className="mb-5 text-gray-800 text-2xl font-semibold">Live Stock Prices</h2>
      {stocks.length === 0 ? (
        <div className="text-center py-10 text-gray-500">
          <p>Loading prices...</p>
        </div>
      ) : (
        <div className="overflow-x-auto">
          <table className="w-full border-collapse text-sm">
            <thead>
              <tr className="bg-gradient-to-r from-indigo-500 to-purple-600 text-white">
                <th className="px-4 py-3 text-left font-semibold text-xs uppercase tracking-wide">Symbol</th>
                <th className="px-4 py-3 text-left font-semibold text-xs uppercase tracking-wide">Price</th>
                <th className="px-4 py-3 text-left font-semibold text-xs uppercase tracking-wide">Change</th>
                <th className="px-4 py-3 text-left font-semibold text-xs uppercase tracking-wide">Change %</th>
              </tr>
            </thead>
            <tbody>
              {stocks.map((stock) => (
                <tr key={stock.symbol} className="border-b border-gray-200 hover:bg-gray-50 transition-colors">
                  <td className="px-4 py-3.5 font-bold text-gray-800">{stock.symbol}</td>
                  <td className="px-4 py-3.5 font-semibold text-gray-900">{formatPrice(stock.price)}</td>
                  <td className={`px-4 py-3.5 font-semibold ${stock.change >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                    {formatChange(stock.change)}
                  </td>
                  <td className={`px-4 py-3.5 font-semibold ${stock.changePct >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                    {formatChangePct(stock.changePct)}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  )
}

export default StockPrices

