// Init DataTable
$(document).ready(function(){
    $(".table-data").DataTable()
    
    
})

var obj = [{}]
// Realtime data when change the quantity
$(document).on('change', '.qty', function(e) {

     //Paksa min max
     var max = parseInt($(this).attr('max'));
     var min = parseInt($(this).attr('min'));
     if ($(this).val() > max) {
        $(this).val(max);
     }
     else if ($(this).val() < min) {
        $(this).val(min);
     } 


    var editTotalPerItem = $("input[name='qty[]']")
              .map(function(){return $(this).val();}).get();
    var title =  $("#table-cart").DataTable().column(2).data()
    var price = $('#table-cart').DataTable().column(4).data()
    var total = 0
    for (let i = 0; i < title.length; i++) {
        let judul = title[i]
        let harga = price[i]
        let qty = parseInt(editTotalPerItem[i])
        total += harga*qty
        obj[0][judul]['qty'] = qty
    }
    console.table(obj)
    $(".price").html("Rp."+total)

    
    
 });

function changeQtyPrice() {
    var title =  $("#table-cart").DataTable().column(2).data()
    var qty = $("#table-cart").DataTable().column(3).data()
    var sum = $('#table-cart').DataTable().column(4).data()
    var tempTotal = 0
    for (let i = 0; i < sum.length; i++) {
        let judul = title[i]
        let price = sum[i]
        let perItem = qty[i]
        tempTotal += perItem * price
        obj[0][judul] = {'qty' : perItem, 'price' : price}
    }

    $(".price").html("Rp."+tempTotal)
    console.table(obj)
}

let apiCart = 'http://localhost:5000/cart/'
function deleteCartItem(idCart) {
    var url = apiCart+idCart
    console.log(url)
    $.ajax({
        url : url,
        type : "DELETE"
    }).done(function (){
        $("#table-cart").DataTable().ajax.reload();

        
        setTimeout(() => {
            changeQtyPrice();
        }, 500);

        
    }).fail(function (response){
        alert("Error")
    })
}

function cartModalShow(){
    var getIdUser = document.getElementById("userIdAndName").getAttribute('data-user-id')

    var url = apiCart + getIdUser
    
    
    $("#table-cart").DataTable({
        responsive:true,
        initComplete : changeQtyPrice,
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

    console.log(dataTable.column(4).data().sum())
    
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
            }else{
                alert('Gagal ditambahkan', 'danger')
            }

            setTimeout(function() {
                bootstrap.Alert.getOrCreateInstance(document.querySelector(".alert")).close();
            }, 1000)
        };
    }
    xhr.send();
}

