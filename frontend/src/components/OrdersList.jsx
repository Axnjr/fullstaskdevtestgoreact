function OrdersList({ orders }) {
  const formatPrice = (price) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2,
      maximumFractionDigits: 2,
    }).format(price)
  }

  const formatDate = (timestamp) => {
    return new Date(timestamp).toLocaleString('en-US', {
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    })
  }

  const ordersArray = Array.isArray(orders) ? orders : []

  // newest first
  const sortedOrders = [...ordersArray].sort((a, b) => 
    new Date(b.timestamp) - new Date(a.timestamp)
  )

  return (
    <div className="bg-white rounded-xl p-6 shadow-md max-h-[calc(100vh-200px)] flex flex-col">
      <h2 className="mb-5 text-gray-800 text-2xl font-semibold">Orders Table</h2>
      {sortedOrders.length === 0 ? (
        <div className="text-center py-10 text-gray-500">
          <p>No orders placed yet</p>
        </div>
      ) : (
        <div className="overflow-y-auto overflow-x-auto flex-1">
          <table className="w-full border-collapse text-sm">
            <thead className="bg-gradient-to-r from-indigo-500 to-purple-600 text-white sticky top-0 z-10">
              <tr>
                <th className="px-2.5 py-3 text-left font-semibold text-xs uppercase tracking-wide">Side</th>
                <th className="px-2.5 py-3 text-left font-semibold text-xs uppercase tracking-wide">Symbol</th>
                <th className="px-2.5 py-3 text-left font-semibold text-xs uppercase tracking-wide">Quantity</th>
                <th className="px-2.5 py-3 text-left font-semibold text-xs uppercase tracking-wide">Price</th>
                <th className="px-2.5 py-3 text-left font-semibold text-xs uppercase tracking-wide">Total</th>
                <th className="px-2.5 py-3 text-left font-semibold text-xs uppercase tracking-wide">Time</th>
              </tr>
            </thead>
            <tbody>
              {sortedOrders.map((order) => (
                <tr key={order.id} className="border-b border-gray-200 hover:bg-gray-50 transition-colors">
                  <td className="px-2.5 py-3 align-middle">
                    <span className={`px-2.5 py-1 rounded text-xs font-bold uppercase tracking-wide inline-block ${
                      order.side === 'buy'
                        ? 'bg-green-100 text-green-600'
                        : 'bg-red-100 text-red-600'
                    }`}>
                      {order.side.toUpperCase()}
                    </span>
                  </td>
                  <td className="px-2.5 py-3 font-semibold text-gray-800">{order.symbol}</td>
                  <td className="px-2.5 py-3 text-gray-600">{order.quantity}</td>
                  <td className="px-2.5 py-3 font-semibold text-gray-800">{formatPrice(order.price)}</td>
                  <td className="px-2.5 py-3 font-bold text-gray-900">
                    {formatPrice(order.quantity * order.price)}
                  </td>
                  <td className="px-2.5 py-3 text-xs text-gray-500">{formatDate(order.timestamp)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  )
}

export default OrdersList

