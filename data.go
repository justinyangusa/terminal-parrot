package main

import "github.com/nsf/termbox-go"

var frames = [...]string{
	`                       N0kxdlccc:cldOX            
                    W0o,..           ,xX          
                   Xo.                 ,xX        
                  K:.     .......       .:0W      
                 K:   ..',;cc::::,.       lN      
                 O' .';:cllllllllc,..     .xW     
                Wd. .',:llloolcccc:,..     'O     
                0' ....';:col:,,,;;;,'.     lN    
               K:......''';cc;,''',;;,.  ...cN    
              K: ':'',,'',;c::::;;:::;'.':,.dW    
            WO;...;,;c:,;:clcccllooc:;,;cc,:K     
           Xo',ol..';:,',;:c::cclllc:,,:c,;kW     
         Nx,'cxkko;.','.',;::c:;;:::;,co;.oN      
       WO:':dkkkkkko;''.,:clol:;;::;,cxo.,K       
     NOc';okkkkkkkkkxc,',;::cc:::;,,cdkl.;X       
    Kc';okkkkkkkkkkkkkl;,,;;:::;;:cdkkkd''O       
   K:.,loooooooollollol:,''',,',:cllolol,.:K      
   O.                                 .   .oN     `,

	`               WKd:..         .:ON                
             WKl.               .:xX              
            WO,     ....          .:0W            
           Wk'   ..';:;;,,;'.       ,0            
           Nc  .',:clllllllc,..     .k            
           Nc .';:cloollllllc;..    .oN           
           K; ...';cool:,,;;:;,'.    ;K           
           k......',cl:,''',,,;;'    .k           
          Nc ..'''',:c:;;;,,,;::,..'''x           
          K, .;::;,:cc:clcclllc:,';c:'d           
          O. .,::,,:cc:cllllol:;,;cc,.oW          
         Wd.':,,'.',;:::::::c:;,;oxOd',0          
        Wk''x0l,'.';cllc;;;;:;,;xXXXXd.;0W        
      W0c';kXX0o,',:cllc::::;,;dKXXXXKd',OW       
    W0c';dKXXXXKx;,;;::::;;;:okKXXXXXXXk,'kW      
   Wk',dKXXXXXXXXklcccccclox0XXXXXXXXXXXx.,0      
   0,.cxxxxxxxxxxxddoooddxxxxxxxxxxxxxxxd,.dW     
   k.  ..................................  :X     `,

	`           NOl,.          .:OW                    
         WO:.               .:kN                  
        Wx.     ....          .c0W                
       Wx.   ..';:;,,,'.        :X                
      Wk.  .';:cllllllc,..      ,0                
      K;  .';:cloolllll:,'.     'O                
     Wd.  ..',:lolc;,;::;,,'.   .xW               
     X:.......';cl:,'',,;;;,.    ;K               
     0''c'.'''';cc:;;,,;;::;'..'''k               
     O.'l;;::,,:cc:ccllllcc;,,:c;;O               
     X:.,;;:;,,ccc:cclolcc:;,;cc,.oN              
      k...,,'.',:::::::c::;,:oxkkc,:xKW           
      0' ..'..':cll:;;;:;;,:xKKKK0kl;,ckN         
      O''lc,'',:clc::::;,,;d0KKKKKKK0kc':OW       
    W0;'dK0Ol,,;;:::;;;:odk0KKKKKKKKKKKx,,kW      
   WO,,xKKKK0dc:::::clxOKKKKKKKKKKKKKKKKx';K      
   K;.:dddddddoollloodddddddddddddddddddo,.xW     
   O.  ....   .                            cN     `,

	`      N0kdlccc::ldOXW                             
   WKd;..          'dX                            
  Xd'.               ,dKW                         
 Xl.    .......        'd0X                       
Nl.  ..';cc::::,.       ..lN                      
0, ..,:cllollllc;..      ..lX                     
0, .',:cooolcccc:;'.    ,ol,:0                    
0,....';cllc;,,;;;,'.   ;xOd.cX                   
K:......,cc;,''',;;;.  .cO0O,'0                   
No'',,'';c:::;;;::::'..;dO0k';K                   
Wd,;c:,;clc:clloolc;,;:cdOKd.cN                   
Wx,;:;',:c::cclllc:,,:loOXXO;'xW                  
 Kl','.',::c:;;:::;,lkOKXXXXOc'cON                
 WO:''.,:lll:;;::;,lOXXXXXXXXKxc;:lkKN            
  WO;.',;:cc:::;,;lOXXXXXXXXXXXX0dc:;:lxXW        
   0'.,,;;:::;;:okKXXXXXXXXXXXXXXXXK0xl;;xN       
   O..:;,,,,,;cdxkkxxkxxxxkkkxkkkxxkkxxd,.cX      
   O. ... ..............................  .dW     `,

	`        WWNNNNNNNW                                
     N0dc;,,,,''';lkN                             
   Xx;.            .,lON                          
 W0:.     ..          .cKW                        
Wk'     .','....        :kK                       
X:  ..',:lllccc;..       .:0W                     
0, .';cclolollll:'.     .,.'xW                    
k. ..';clool:::::;'.    ,ol''O                    
c .....,:ll:,'',;;,,.   ,od:.dW                   
: ....'.,cc;,,'',;;;.  .cxx;.x                    
l..,;;,,:cc::c::cc::'.,:oxd''O                    
k'';c:,;clc:clloolc;,;ccdxxc.cK                   
K:.,;,',;:::cccccc;,,cldkOOkl',dKW                
Nc..''.';:cc:;;;::;;lxkkOOOOOxc,,:okKNW           
Wk..''.,clllc:::;,,ckOOOOOOOOOOkdc;,,;cdKW        
 No..,',;::c:::;,;lkOOOOOOOOOOOOOOOkxoc,'oX       
  Nl..'',;;;;''';clooooooooooooooooooool,.;0W     
   k.  ......                              cN     `,

	`                                                  
                                                  
         WKkxxoooloxOKN                           
       Xx:..         .;o0NW                       
     Nx,.               .;dN                      
    Xc.    ......         .lX                     
   Xl.  ...;:c:;:;'.       .lKW                   
  Nl. ..,:cllollllc,.       .,oKW                 
  x.  .';:looollccc:,.      ,c',OW                
 Wd. ....':lol:;,,;;;,..    ;xl.:X                
  0, .....';cc;,'',,;;,.   .ckl.;X                
  Nc..',,'';c:::;,;;::;....:dkc.oW                
  Nc.':c:,;cl:clllllc:;,,::ldkc.lN                
  X:.';:;',cc::cclllc:,,:lloxOd,.oKW              
  K,.,,,'.';::c::;:::;,cdxxkkOkxl,,ckN            
  K,.l:'..,clll:;;::;,cxOOOOOkkOOko;,;lkXW        
  Nc.;:,.';:ccc:::;,',clooooooooooolc;...:kN      
   k.   .',;;::::;'..                     .oN     `,

	`                                                  
                                                  
                                                  
               WNKOxxxddxxxOKW                    
            WXkl,..         .;xX                  
          WKo,.                ;0W                
         Xo'      ...           ,kW               
       Wk,      .';;,''..        'xN              
      Wx'.  ..',:llllllc,..       .lX             
      0,.,...,:cloooolllc:;..    ...oN            
      k.,:...',:loolc;;:::;,'.   .c,'O            
      O'':.....',clc;'',,;;;;'.  .l:.dW           
      k..;,..'''':c:;,'',,;::,. .;d:.oW           
     No...',;:;,;cc:ccccclcc:,'':l:.;0            
    Wk',o:',:c;,:lc:cllooolc:,,:c,.cK             
   WO,'dOkl;,,'',::::c:cccc:,,coo;.;odk0XW        
   Nc.:oddo;''..;:clc:;;:::;,;lool:,'....;xN      
   X;  ......''':looc::::;;,..........    .dW     `,

	`                                                  
                                                  
                     WX0kxdoooolodOKW             
                   NOl,.           .c0W           
                 W0c.                .c0W         
                Wk,     .......        .dN        
                O'   ..':cc::;:;.       .kW       
              W0;  .';:cllllllll:'.      lW       
             Nx,. ..,;cloooolcccc;'.     lW       
            Nd';;....',:lool:,,;;;,,.   .oW       
            k',dl....'.':ll:;,'',,;;,.  .lN       
           Xc.:xd;.',,',:cc::;;;;;::,..';lK       
         W0:..,dd:,:c;,:clc:clllolc:,,;ccdX       
       W0l',l:.;o:,::,';ccc::cclll:;,,cccOW       
     WKl',cdxxl'';;,'..,;:::::;;:::,;lo,;0        
    Nd,'cdxxkkxd:'..'.';clllc:;;::,;lxd,.xW       
   No..clllllllllc,...';:cccc:::;,';looc..dX      
   X;                .',;;::::;,'.......  .dW     `,

	`                                                  
                          WNK0000000KNW           
                      WXxol,.........'l0W         
                     Xd,.              .:kN       
                   WO;.     ....         .:0W     
                  WO,     ..,;;,,,'.       ;K     
                  X:  ..,,;cllllllc,.      .k     
                Nk;. .';:cclloollllc;.     .x     
              W0:.,'...',,:lool:;;;:;,'.   .lN    
             Wk'.:d:.....'',cl:,'',,;,;'   .lN    
            Wx. .lxl'.''''',:c:;,,,,;::,...':0    
            0,...:xo;,::;,,;cc:cccclcc:,',c:,dW   
          W0;.:l'.co;;::,',:lc:clllolc;,;cc'.dW   
        W0l',ldxo,.,;,,,..',:::::::c:;,,co:.;K    
      WOc',cdxxxxdc,..''..';cllc;;;::;,cdo''kW    
    WOc',ldxxxxxxxxdl:;''.,:lllc:::;,,cdd;.oW     
   Wx..;clllllllllllllc,''';;::::;,'';clc..x      
   X:   .             . ..';;::;,'.        lN     `,

	`                             WNNNNNNNNW           
                        WWXko:,,,,'''';o0W        
                      WKdc,.            .c0N      
                     Wk,       ..         .c0W    
                    Nd.     ..,,'....       ,O    
                   Wk'  ...';clllccc,.      .dW   
                  Xo.  .,;:clllllollc;..    .oW   
                W0;.'. .'',:coolc::::;,'.    cX   
               Wk'.:c......',clc;,',;;;,'.   .dW  
              Wx. .lo,....''';c:;,''',;;,. ...'k  
              O'...co:',;:;,,:c::c::ccc:;'.;c,.cN 
            W0;.::.'l:,:cc:,;ll:ccloolc:,,:cc,.:N 
          WXd'.cdd:..,,;;,'',::::ccccc:;,;clc'.xW 
        WKo,';lddddl,...''..,:ccc:;;::;,;ldo,.oN  
      W0c'':oddddddddl:,'''';clol::::;,,cdo,.oN   
    W0c',coddddddddddddoc;,',::cc::;,,;ldd:.:X    
   Wk'.,:cccccccccccccccc:,',;;;;,'.';:cc:..kW    
   X:                       ......         ,K     `,
}

var colors = [...]termbox.Attribute{
	termbox.Attribute(425),
	termbox.Attribute(227),
	termbox.Attribute(47),
	termbox.Attribute(263),
	termbox.ColorBlue,
	termbox.Attribute(275),
	termbox.Attribute(383),
	termbox.Attribute(419),
	termbox.Attribute(202),
	termbox.Attribute(204),
}
