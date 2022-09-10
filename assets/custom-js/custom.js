var tableData =  $(".table-data")
var tableCart = $("#table-cart")

// Init DataTable
$(document).ready(function(){
    tableData.DataTable()
})

var arrObj = []

// checkout function
$(".modal-content").on('click','#btnCheckout',function(){
    if(confirm("Apakah anda yakin ingin memesan ini ?") == true){
        checkoutAndChangeQty()
        url = "http://localhost:5000/checkout/"
        if (arrObj.length > 0){
            $.ajax({
                url : url,
                type : "POST",
                data : JSON.stringify(arrObj),
                contentType : "application/json"
                
            }).done(function (){
                // Hapus semua item cart dari Backend dan frontend
                var idCart =  tableCart.DataTable().column(0).data()
                for (let i = 0; i < idCart.length; i++) {
                    let el = idCart[i];
                    deleteCartItem(el)
                }
                arrObj.splice(0,arrObj.length)
            }).fail(function (){
                alert("Error")
            })
        }else if (arrObj.length === 0){
            alert("Tidak ada barang yang di checkout")
        }else if (arrObj.length === 4){
            alert("Barang tidak boleh lebih dari 4")
        }
    }
   
    
})

// Realtime data when change the quantity
$(".modal-content").on('change', '.qty', function(e) {
     //Paksa min max
     var max = parseInt($(this).attr('max'));
     var min = parseInt($(this).attr('min'));
     if ($(this).val() > max) {
        $(this).val(max);
     }
     else if ($(this).val() < min) {
        $(this).val(min);
     } 

    checkoutAndChangeQty()
 });


// function changeQtyDB(params) {
    
// }
 
function checkoutAndChangeQty(){
    var editTotalPerItem = $("input[name='qty[]']")
              .map(function(){return $(this).val();}).get();
    var idCart =  tableCart.DataTable().column(0).data()
    var price = $('#table-cart').DataTable().column(4).data()
    var total = 0
    for (let i = 0; i < idCart.length; i++) {
        let cartId = idCart[i]
        let harga = price[i]
        let qty = parseInt(editTotalPerItem[i])
        total += harga*qty
        let index = arrObj.findIndex(key => key.idCart === cartId);
        if (index == -1){
            arrObj.push({
                idCart : cartId,
                qty : qty,
                harga : harga
            })
        }else{
            arrObj[i].idCart = cartId
            arrObj[i].qty = qty
            arrObj[i].harga = harga
        }
    }
    console.table(arrObj)
    $(".price").html("Rp."+total)
}


function searchKeyForDelete(idCart) {
    var __FOUND = -1;
    for(var i=0; i<arrObj.length; i++) {
        if(arrObj[i].idCart == idCart) {
            __FOUND = i;
            break;
        }
    }
    arrObj.splice(__FOUND,1)
}

let apiCart = 'http://localhost:5000/cart/'
function deleteCartItem(idCart) {
    var url = apiCart+idCart
    $.ajax({
        url : url,
        type : "DELETE"
    }).done(function (){
        tableCart.DataTable().ajax.reload();
        // delete arrObj[0][idCart]
        searchKeyForDelete(idCart) 
        setTimeout(() => {
            checkoutAndChangeQty();
        }, 1000);

        
    }).fail(function (response){
        alert("Error")
    })
}

function cartModalShow(){
    var getIdUser = document.getElementById("userIdAndName").getAttribute('data-user-id')

    var url = apiCart + getIdUser
    
    
    tableCart.DataTable({
        responsive:true,
        initComplete : checkoutAndChangeQty,
        bDestroy: true,
        ajax : {
            url : url,
            dataSrc : '',
            type : 'POST'
        },
        columnDefs: [ 
            {
                targets:  0 ,
                data: 'IdCart',
                render : function (data, type, full, meta) {
                    return '<span data-id='+data+'>'+data+'</span>';
                }
            },
            {
                targets:  1 ,
                data: 'CoverBuku',
                orderable : false,
                render : function (data, type, full, meta) {
                    return '<img src="/images/'+data+'" width="75px" height="75px"/>';
                }
            },
            {
                targets:  2 ,
                data: 'JudulBuku',
              
            },
            {
                targets:  3 ,
                data: 'TotalPerItem',
                orderable : false,
                render : function (data, type, full, meta) {
                    return '<input type="number" min=0 max=4 class="form-control text-center qty" name="qty[]" style="max-width: 5rem;" value="'+data+'"/>';
                }

            },
            {
                targets:  4 ,
                data: 'Harga',
            },
            {
                targets : 5,
                data: 'IdCart',
                orderable : false,
                render : function(data,type,full,meta){
                    return '<button class="btn btn-danger btn-sm" onclick="deleteCartItem('+data+');"><i class="fa fa-times"></i></button>'
                }
            }
        ]
        
        ,
    })

    var cartModal = new bootstrap.Modal(document.getElementById('cartModal'))
    cartModal.show();

    
}
// Get Details of Book with Button (Ajax)
function BookDetailAndCartFunc(book) {

    var bookId = book.getAttribute("data-book-id")
    var whichButton = book.getAttribute("data-button")
    // console.log(whichButton)
   
    const xhr = new XMLHttpRequest();
    const url = 'http://localhost:5000/books/'+bookId

    if (whichButton == "detail"){
        xhr.open("GET", url, true);
        xhr.onload = function (data) {              
            dataParsed = JSON.parse(data.target.response)   
            document.getElementById("modalTitle").innerHTML = dataParsed[0].Judul
            document.getElementById("deskripsi").innerText = dataParsed[0].Deskripsi
            document.getElementById("kategori").innerText = dataParsed[0].Kategori
            document.getElementById("pengarang").innerText = dataParsed[0].Pengarang
            document.getElementById("penerbit").innerText = dataParsed[0].Penerbit
            document.getElementById("tahun").innerText = dataParsed[0].Tahun
            document.getElementById("stok").innerText = dataParsed[0].Stok
        };
    }else if (whichButton == "cart"){
        var getIdUser = document.getElementById("userIdAndName").getAttribute('data-user-id')
        xhr.open("POST",url+"/"+getIdUser,true)
        xhr.onload = function (response) {
            const alertPlaceholder = document.getElementById('liveAlertPlaceholder')
            const alert = (message, type) => {
            const wrapper = document.createElement('div')
            wrapper.innerHTML = [
                `<div class="alert alert-${type} alert-dismissible fade show" role="alert">`,
                `   <div>${message}</div>`,
                '   <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>',
                '</div>'
            ].join('')

            alertPlaceholder.append(wrapper)
            }

            if(response.target.status == 200 && response.target.readyState == 4){
                alert('Sukses ditambahkan', 'success')
            }else if(response.target.status == 406){
                alert('Gagal menambahkan karena melebihi limit', 'danger')
            }

            setTimeout(function() {
                bootstrap.Alert.getOrCreateInstance(document.querySelector(".alert")).close();
            }, 1000)
        };
    }
    xhr.send();
}
