# PowerShell script to test request ID functionality
Write-Host "Testing Request ID Functionality..." -ForegroundColor Green

# Start server in background
Write-Host "Starting server..." -ForegroundColor Yellow
$job = Start-Job -ScriptBlock {
    Set-Location "d:\KERJA\BelajarGO\belajar-golang-rest"
    go run .
}

# Wait for server to start
Start-Sleep -Seconds 3

try {
    # Test 1: Valid request - should show request ID in logs
    Write-Host "`nTest 1: Getting all categories (should work)" -ForegroundColor Cyan
    $response1 = curl.exe -s -X GET "http://localhost:3000/api/categories" -H "X-API-Key: RAHASIA"
    Write-Host "Response: $response1"
    
    # Test 2: Invalid category ID - should show request ID in error logs  
    Write-Host "`nTest 2: Getting invalid category (should trigger error with request ID)" -ForegroundColor Cyan
    $response2 = curl.exe -s -X GET "http://localhost:3000/api/categories/999" -H "X-API-Key: RAHASIA"  
    Write-Host "Response: $response2"
    
    # Test 3: Update invalid category - should show request ID in error logs
    Write-Host "`nTest 3: Updating invalid category (should trigger error with request ID)" -ForegroundColor Cyan  
    $response3 = curl.exe -s -X PUT "http://localhost:3000/api/categories/999" -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -d '{"name":"Updated Category"}'
    Write-Host "Response: $response3"
    
} catch {
    Write-Host "Error occurred: $_" -ForegroundColor Red
} finally {
    # Stop server
    Write-Host "`nStopping server..." -ForegroundColor Yellow
    Stop-Job $job -ErrorAction SilentlyContinue
    Remove-Job $job -ErrorAction SilentlyContinue
}

Write-Host "`nTest completed. Check console output for request ID logs." -ForegroundColor Green