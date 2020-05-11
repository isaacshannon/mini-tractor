difference(){
    union(){
        difference(){
            cylinder(r=38, h=10);
            translate([0,0,-1])cylinder(r=35, h=30);
        }
        translate([-12,11,0])cube([24,24,10]);
        translate([0,6,4])cube([71,2,10], center=true);
        translate([0,8,4])cube([2,6,10], center=true);
    };
    translate([-10,13,-1])cube([20,20,60]);
    translate([0,-40,-1])cube([70,50,60], center=true);

};